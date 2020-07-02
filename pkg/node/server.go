package node

import (
	"context"
	"github.com/elliotcourant/pubby-mcsubface/pkg/protos"
)

var (
	_ protos.RaftServer = &Server{}
)

type Server struct {
	node      *Node
	transport *Transport
}

func (s *Server) AppendEntries(ctx context.Context, request *protos.AppendEntriesRequest) (*protos.AppendEntriesResponse, error) {
	panic("implement me")
}

func (s *Server) RequestVote(ctx context.Context, request *protos.RequestVoteRequest) (*protos.RequestVoteResponse, error) {
	panic("implement me")
}

func (s *Server) InstallSnapshot(ctx context.Context, request *protos.InstallSnapshotRequest) (*protos.InstallSnapshotResponse, error) {
	panic("implement me")
}

func (s *Server) TimeoutNow(ctx context.Context, request *protos.TimeoutNowRequest) (*protos.TimeoutNowResponse, error) {
	panic("implement me")
}
