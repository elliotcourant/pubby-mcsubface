package transport

import (
	"github.com/hashicorp/raft"
	"io"
)

var (
	_ raft.Transport = &Transport{}
)

type Transport struct {
	connectionPool

	consumeChannel chan raft.RPC
}

func (t *Transport) Consumer() <-chan raft.RPC {
	panic("implement me")
}

func (t *Transport) LocalAddr() raft.ServerAddress {
	panic("implement me")
}

func (t *Transport) AppendEntriesPipeline(id raft.ServerID, target raft.ServerAddress) (raft.AppendPipeline, error) {
	panic("implement me")
}

func (t *Transport) AppendEntries(id raft.ServerID, target raft.ServerAddress, args *raft.AppendEntriesRequest, resp *raft.AppendEntriesResponse) error {
	panic("implement me")
}

func (t *Transport) RequestVote(id raft.ServerID, target raft.ServerAddress, args *raft.RequestVoteRequest, resp *raft.RequestVoteResponse) error {
	panic("implement me")
}

func (t *Transport) InstallSnapshot(id raft.ServerID, target raft.ServerAddress, args *raft.InstallSnapshotRequest, resp *raft.InstallSnapshotResponse, data io.Reader) error {
	panic("implement me")
}

func (t *Transport) EncodePeer(id raft.ServerID, addr raft.ServerAddress) []byte {
	panic("implement me")
}

func (t *Transport) DecodePeer(bytes []byte) raft.ServerAddress {
	panic("implement me")
}

func (t *Transport) SetHeartbeatHandler(cb func(rpc raft.RPC)) {
	panic("implement me")
}

func (t *Transport) TimeoutNow(id raft.ServerID, target raft.ServerAddress, args *raft.TimeoutNowRequest, resp *raft.TimeoutNowResponse) error {
	panic("implement me")
}
