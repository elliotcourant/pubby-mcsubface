package node

import (
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/memberlist"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		// The network addresses of at least one other node in the cluster.
		Peers []string

		// Name is the name that this node will use to identify itself to the cluster. If this is left blank it will be set
		// to a UUID.
		Name string

		// Log for logging things.
		Log *logrus.Entry
	}

	// Node represents a single worker in a cluster. This is used to address others in the cluster and coordinate with
	// them but serves no other real purpose. This cluster uses the gossip protocol so the consistency is not perfect.
	Node struct {
		cluster *memberlist.Memberlist
		config  Config
		log     *logrus.Entry
	}
)

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

	// If there were any peers specified in the config then try to join those peers within the cluster.
	if len(config.Peers) > 0 {
		if _, err = cluster.Join(config.Peers); err != nil {
			return nil, errors.Wrap(err, "failed to join cluster")
		}
	}

	node := &Node{
		cluster: cluster,
		config:  *config,
		log:     config.Log,
	}

	return node, nil
}

func (n *Node) Close() error {
	// If we defer this then we can make sure that this node will at least get shutdown properly when trying to leave the
	// cluster. This will happen even if the node times out trying to leave the cluster gracefully.
	defer n.cluster.Shutdown()

	// Try to leave the cluster gracefully. If it fails then wrap the error.
	return errors.Wrap(n.cluster.Leave(30*time.Second), "failed to gracefully leave the cluster")
}
