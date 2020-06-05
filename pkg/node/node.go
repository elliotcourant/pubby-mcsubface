package node

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/memberlist"
	"github.com/pkg/errors"
)

type (
	Config struct {
		// The network addresses of at least one other node in the cluster.
		Peers []string

		// Name is the name that this node will use to identify itself to the cluster. If this is left blank it will be set
		// to a UUID.
		Name string
	}
)

type Node struct {
	config  Config
	cluster *memberlist.Memberlist
}

func NewNode(config *Config) (*Node, error) {
	// If the name is not populated then make one.
	if len(config.Name) == 0 {
		name, err := uuid.GenerateUUID()
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate a unique node name")
		}

		config.Name = name
	}

	// We just need to change a few of the parameters, but we can use the default config provided for the most part.
	memberlistConfig := memberlist.DefaultLANConfig()
	memberlistConfig.Name = config.Name

	cluster, err := memberlist.Create(memberlistConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cluster")
	}

	node := &Node{
		config:  *config,
		cluster: cluster,
	}

	return node, nil
}
