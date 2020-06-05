package testutils

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListeners(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		// Create 3 listeners
		size := 3

		listeners := NewListeners(t, size)
		assert.Len(t, listeners, size, "number of listeners does not match expected")

		// Make sure that all the addresses can be properly resolved.
		for _, listener := range listeners {
			addr := listener.Addr().String()
			_, err := net.ResolveTCPAddr("tcp", addr)
			assert.NoError(t, err, "address should have resolved normally: %s", addr)
		}

		// Close all of the listeners.
		for _, listener := range listeners {
			err := listener.Close()
			assert.NoError(t, err, "listener should have closed successfully")
		}
	})
}

func TestGetLocalAddress(t *testing.T) {
	address := GetLocalAddress(t)
	assert.NotEmpty(t, address, "address should not be empty")
}

func TestNewAddresses(t *testing.T) {
	t.Run("small", func(t *testing.T) {
		size := 3

		addresses, strings := NewAddresses(t, size)
		assert.Len(t, addresses, size, "number of addresses does not match expected")
		assert.Len(t, strings, size, "number of strings does not match expected")
	})

	t.Run("big", func(t *testing.T) {
		size := 16

		addresses, strings := NewAddresses(t, size)
		assert.Len(t, addresses, size, "number of addresses does not match expected")
		assert.Len(t, strings, size, "number of strings does not match expected")
	})
}
