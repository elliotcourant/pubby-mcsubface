package node

import (
	"github.com/hashicorp/raft"
	"net"
)

type (
	Config struct {
		Peers         []string
		ListenAddress string
	}

	Node struct {
		config     Config
		raftServer *raft.Server
	}
)

func NewNode(config Config) (*Node, error) {

}

func setupListener(config Config) (net.Listener, error) {

}
