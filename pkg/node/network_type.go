package node

//go:generate stringer -type=NetworkType -output=network_type.strings.go
type NetworkType uint8

const (
	Unknown NetworkType = iota
	Udp
	Tcp
	Unix
)
