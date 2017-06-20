package jsnet

// An Addr identifies a connection
type Addr struct {
	id string
}

// Network returns jsnet
func (addr Addr) Network() string {
	return "jsnet"
}

// String returns the address
func (addr Addr) String() string {
	return addr.id
}
