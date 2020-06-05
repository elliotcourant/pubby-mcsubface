package node

import (
	"testing"
	"time"

	"github.com/elliotcourant/pubby-mcsubface/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNode(t *testing.T) {
	t.Run("single node", func(t *testing.T) {
		log := testutils.NewLog(t)

		config := Config{
			Peers: nil,
			Log:   log,
		}

		node, err := NewNode(&config)
		assert.NoError(t, err)
		assert.NotNil(t, node)

		err = node.Close()
		assert.NoError(t, err, "should close gracefully")
	})

	t.Run("a few nodes", func(t *testing.T) {
		log := testutils.NewLog(t)
		size := 3

		addresses, addressStrings := testutils.NewAddresses(t, size)

		nodes := make([]*Node, size)

		for i := 0; i < size; i++ {
			config := Config{
				Peers:         addressStrings,
				Log:           log,
				ListenAddress: addresses[i].IP.String(),
				ListenPort:    addresses[i].Port,
			}

			node, err := NewNode(&config)
			require.NoError(t, err)

			nodes[i] = node
		}

		// Let the cluster run for a few seconds.
		time.Sleep(3 * time.Second)

		// Kill all the nodes in the cluster.
		for _, node := range nodes {
			err := node.Close()
			assert.NoError(t, err, "node should have closed successfully")
		}
	})
}
