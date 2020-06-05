package node

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	t.Run("single node", func(t *testing.T) {
		log := logrus.New()
		entry := logrus.NewEntry(log)

		config := Config{
			Peers: nil,
			Log:   entry,
		}

		node, err := NewNode(&config)
		assert.NoError(t, err)
		assert.NotNil(t, node)

		err = node.Close()
		assert.NoError(t, err, "should close gracefully")
	})
}
