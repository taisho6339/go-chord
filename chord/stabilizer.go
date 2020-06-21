package chord

import (
	"context"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

type Stabilizer interface {
	Stabilize(ctx context.Context)
}

type SuccessorStabilizer struct {
	Node *LocalNode
}

func (s SuccessorStabilizer) Stabilize(ctx context.Context) {
	// Check whether there are other nodes between s and the successor
	n, err := s.Node.nodeRepo.PredecessorRPC(ctx, s.Node.Successor)
	if err != nil {
		log.Warnf("successor stabilizer failed. err = %#v", err)
		return
	}
	if n != nil && n.ID.Between(s.Node.ID, s.Node.Successor.ID) {
		s.Node.Successor = n
	}
	if err := s.Node.nodeRepo.NotifyRPC(ctx, &s.Node.NodeRef, s.Node.Successor); err != nil {
		log.Warnf("successor stabilizer notify failed. err = %#v", err)
	}
}

type FingerTableStabilizer struct {
	Node *LocalNode
}

func (s FingerTableStabilizer) Stabilize(ctx context.Context) {
	n := rand.Intn(bitSize-2) + 2 // [2,m)
	succ, err := s.Node.FindSuccessor(ctx, s.Node.FingerTable[n].ID)
	if err != nil {
		log.Warnf("finger table stabilizer failed. err = %#v, finger id = %x", err, s.Node.FingerTable[n].ID)
		return
	}
	s.Node.FingerTable[n].Node = succ
}
