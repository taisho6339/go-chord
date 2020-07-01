package chord

import (
	"context"
	"fmt"
	"github.com/taisho6339/gord/model"
	"sync"
)

type exclusiveNodeList struct {
	nodes   []RingNode
	hostMap map[string]struct{}
	lock    sync.Mutex
}

func newNodeList(cap int) *exclusiveNodeList {
	return &exclusiveNodeList{
		nodes:   emptyNodes(cap),
		hostMap: map[string]struct{}{},
	}
}

func (q *exclusiveNodeList) refreshHostMap() {
	hostMap := map[string]struct{}{}
	for _, node := range q.nodes {
		hostMap[node.Reference().Host] = struct{}{}
	}
	q.hostMap = hostMap
}

func (q *exclusiveNodeList) hasHostKey(host string) bool {
	_, ok := q.hostMap[host]
	return ok
}

func (q *exclusiveNodeList) head() (RingNode, error) {
	if len(q.nodes) == 0 {
		return nil, ErrNoSuccessorAlive
	}
	return q.nodes[0], nil
}

func emptyNodes(cap int) []RingNode {
	return make([]RingNode, 0, cap)
}

func (q *exclusiveNodeList) appendHead(node RingNode) {
	if node == nil {
		return
	}
	if q.hasHostKey(node.Reference().Host) {
		return
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.hostMap[node.Reference().Host] = struct{}{}
	newNodes := append(emptyNodes(cap(q.nodes)), node)
	if len(q.nodes) >= cap(q.nodes) {
		q.nodes = append(newNodes, q.nodes[:len(q.nodes)-1]...)
		return
	}
	q.nodes = append(newNodes, q.nodes[:]...)
}

func (q *exclusiveNodeList) join(offset int, nodes []RingNode) {
	if len(nodes) == 0 {
		return
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(nodes) > (cap(q.nodes) - offset) {
		nodes = nodes[:cap(q.nodes)-offset]
	}

	q.nodes = q.nodes[0:offset]
	q.refreshHostMap()
	for _, node := range nodes {
		if q.hasHostKey(node.Reference().Host) {
			continue
		}
		q.nodes = append(q.nodes, node)
	}
	q.refreshHostMap()
}

type LocalNode struct {
	*model.NodeRef

	fingerTable []*Finger
	successors  *exclusiveNodeList
	predecessor RingNode
	isShutdown  bool
}

func NewLocalNode(host string) *LocalNode {
	id := model.NewHashID(host)
	return &LocalNode{
		NodeRef:     model.NewNodeRef(host),
		fingerTable: NewFingerTable(id),
	}
}

func (l *LocalNode) initSuccessors(suc RingNode) {
	l.successors = newNodeList(l.ID.Size() / 2)
	l.setSuccessor(suc)
}

func (l *LocalNode) setSuccessor(suc RingNode) {
	l.successors.appendHead(suc)
	l.fingerTable[0].Node = suc
}

func (l *LocalNode) Shutdown() {
	l.isShutdown = true
}

func (l *LocalNode) CreateRing() {
	l.initSuccessors(l)
	l.predecessor = l
	for _, finger := range l.fingerTable {
		finger.Node = l
	}
}

func (l *LocalNode) JoinRing(ctx context.Context, existNode RingNode) error {
	successor, err := existNode.FindSuccessorByTable(ctx, l.ID)
	if err != nil {
		return fmt.Errorf("find successor failed. err = %#v", err)
	}
	l.initSuccessors(successor)
	firstSuc, err := l.successors.head()
	if err != nil {
		return err
	}
	err = firstSuc.Notify(ctx, l)
	if err != nil {
		return fmt.Errorf("notify failed. err = %#v", err)
	}
	successors, err := firstSuc.GetSuccessors(ctx)
	if err != nil {
		return fmt.Errorf("get successors failed. err = %#v", err)
	}
	l.successors.join(1, successors)
	return nil
}

func (l *LocalNode) Ping(_ context.Context) error {
	if l.isShutdown {
		return ErrNodeUnavailable
	}
	return nil
}

func (l *LocalNode) Reference() *model.NodeRef {
	return l.NodeRef
}

func (l *LocalNode) GetSuccessors(_ context.Context) ([]RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	return l.successors.nodes, nil
}

func (l *LocalNode) GetPredecessor(_ context.Context) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	return l.predecessor, nil
}

func (l *LocalNode) FindSuccessorByList(ctx context.Context, id model.HashID) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	succ, err := l.successors.head()
	if err != nil {
		return nil, err
	}
	if l.ID.Equals(succ.Reference().ID) {
		return l, nil
	}
	if id.Equals(l.ID) {
		return l, nil
	}
	if id.Between(l.ID, succ.Reference().ID) {
		return succ, nil
	}
	return succ.FindSuccessorByList(ctx, id)
}

func (l *LocalNode) FindSuccessorByTable(ctx context.Context, id model.HashID) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	node, err := l.findPredecessor(ctx, id)
	if err != nil {
		return l.FindSuccessorByList(ctx, id)
	}
	successors, err := node.GetSuccessors(ctx)
	if err != nil {
		return nil, err
	}
	for _, successor := range successors {
		if err := successor.Ping(ctx); err == nil {
			return successor, nil
		}
	}
	return nil, ErrNoSuccessorAlive
}

func (l *LocalNode) findPredecessor(ctx context.Context, id model.HashID) (RingNode, error) {
	var (
		targetNode RingNode = l
	)
	for {
		successor, err := targetNode.GetSuccessors(ctx)
		if err != nil {
			return nil, err
		}
		if successor == nil || len(successor) <= 0 {
			return nil, ErrNotFound
		}
		suc := successor[0]
		if targetNode.Reference().ID.Equals(suc.Reference().ID) {
			return targetNode, nil
		}
		if id.Between(targetNode.Reference().ID, suc.Reference().ID.Add(1)) {
			break
		}
		node, err := targetNode.FindClosestPrecedingNode(ctx, id)
		if err != nil {
			return nil, ErrNotFound
		}
		targetNode = node
	}
	return targetNode, nil
}

func (l *LocalNode) FindClosestPrecedingNode(_ context.Context, id model.HashID) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	for i := range l.fingerTable {
		finger := l.fingerTable[len(l.fingerTable)-(i+1)]
		// If the fingerTable has not been updated
		if finger.Node == nil {
			return nil, ErrStabilizeNotCompleted
		}
		if finger.Node.Reference().ID.Between(l.ID, id) {
			return finger.Node, nil
		}
	}
	return l, nil
}

func (l *LocalNode) Notify(_ context.Context, node RingNode) error {
	if l.isShutdown {
		return ErrNodeUnavailable
	}
	if l.predecessor == nil || node.Reference().ID.Between(l.predecessor.Reference().ID, l.ID) {
		l.predecessor = node
	}
	return nil
}
