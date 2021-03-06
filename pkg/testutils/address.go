package testutils

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

// NewListeners will create N net.Listeners. If a listener cannot be created then this test fill fail.
func NewListeners(t *testing.T, n int) []net.Listener {
	listeners := make([]net.Listener, n)
	for i := 0; i < n; i++ {
		listener, err := net.Listen("tcp", ":")
		require.NoError(t, err, "failed to create listener")
		listeners[i] = listener
	}

	return listeners
}

func GetListenerAddresses(t *testing.T, listeners ...net.Listener) []string {
	addresses := make([]string, len(listeners))
	for i, listener := range listeners {
		addresses[i] = listener.Addr().String()
	}

	return addresses
}

func GetLocalAddress(t *testing.T) string {
	addrs, err := net.InterfaceAddrs()
	require.NoError(t, err, "must be able to list interfaces")

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	// If we could not arrive at an IP address then we can default to localhost as a loopback address. This is only for
	// testing anyway so this should be fine.
	return "localhost"
}

func NewAddresses(t *testing.T, n int) ([]net.TCPAddr, []string) {
	localAddress := GetLocalAddress(t)
	addresses := make([]net.TCPAddr, 0, n)
	addressStrings := make([]string, 0, n)

	for i := 0; i < n; i++ {
		a, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:", localAddress))
		require.NoError(t, err, "should be able to resolve address")

		l, err := net.ListenTCP("tcp", a)
		require.NoError(t, err, "should be able to allocate port")

		addr, ok := l.Addr().(*net.TCPAddr)
		require.True(t, ok, "should be a TCPAddr")
		require.NotNil(t, addr, "TCPAddr should not be nil")

		// Make sure that we have not already seen this address.
		require.NotContains(t, addresses, addr.String(), "address must be unique")
		addressStrings = append(addressStrings, addr.String())
		addresses = append(addresses, *addr)

		err = l.Close()
		require.NoError(t, err, "allocation should close successfully")
	}

	return addresses, addressStrings
}
