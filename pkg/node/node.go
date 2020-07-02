package node

import (
	"github.com/elliotcourant/pubby-mcsubface/pkg/protos"
	"github.com/hashicorp/raft"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type (
	Config struct {
		Peers             []string
		NetworkType       NetworkType
		ListenAddress     string
		ConnectionTimeout time.Duration
	}

	Node struct {
		config     Config
		raftServer *raft.Server
		listener   net.Listener
		grpcServer *grpc.Server
		server     *Server
	}
)

func NewNode(config Config) (*Node, error) {
	listener, err := setupListener(config)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(config.ConnectionTimeout),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     0,
			MaxConnectionAge:      0,
			MaxConnectionAgeGrace: 0,
			Time:                  0,
			Timeout:               1 * time.Second,
		}),
	)

	node := &Node{
		config:     config,
		raftServer: nil,
		listener:   listener,
		grpcServer: grpcServer,
	}

	server := newServer(node)

	protos.RegisterRaftServer(grpcServer, server)

	return nil, nil
}

func setupListener(config Config) (net.Listener, error) {
	var netType string
	switch config.NetworkType {
	case Tcp:
		netType = "tcp"
	case Unix:
		netType = "unix"
	case Udp:
		fallthrough
	default:
		netType = "udp"
	}

	listener, err := net.Listen(netType, config.ListenAddress)
	if err != nil {
		return nil, errors.Wrapf(err,
			"failed to create listener for network/address: %s/%s", netType, config.ListenAddress,
		)
	}

	return listener, nil
}
