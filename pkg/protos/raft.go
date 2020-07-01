package protos

import (
	"github.com/hashicorp/raft"
)

func InstallSnapshotRequestToProto(request *raft.InstallSnapshotRequest, data []byte) *InstallSnapshotRequest {
	return &InstallSnapshotRequest{
		RPCHeader: &RPCHeader{
			ProtocolVersion: ProtocolVersion(request.RPCHeader.ProtocolVersion),
		},
		SnapshotVersion:    SnapshotVersion(request.SnapshotVersion),
		Term:               request.Term,
		Leader:             request.Leader,
		LastLogIndex:       request.LastLogIndex,
		LastLogTerm:        request.LastLogTerm,
		Peers:              request.Peers,
		Configuration:      request.Configuration,
		ConfigurationIndex: request.ConfigurationIndex,
		Data:               data,
	}
}
