package jsnet

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
)

type (
	// A Host identifies a host
	Host [32]byte
)

// HostFromString creates a host from a string
func HostFromString(host string) Host {
	var h Host
	bs, err := hex.DecodeString(host)
	if err != nil {
		panic(err)
	}
	copy(h[:], bs)
	return h
}

// String returns the host as a hex-encoded string
func (h Host) String() string {
	return hex.EncodeToString(h[:])
}

type (
	// An Addr identifies an address in the network
	Addr struct {
		Host Host
		Port uint16
	}
)

func AddrFromString(hostport string) Addr {
	host, sport, _ := net.SplitHostPort(hostport)
	port, _ := strconv.Atoi(sport)
	return Addr{
		Host: HostFromString(host),
		Port: uint16(port),
	}
}

// Network returns jsnet
func (addr Addr) Network() string {
	return "jsnet"
}

// String returns the address
func (addr Addr) String() string {
	return fmt.Sprintf("%s:%d", addr.Host, addr.Port)
}
