package jsnet

type (
	// PacketType identifies the type of data in the packet
	PacketType byte

	// A Packet is a packet of data in the network
	Packet struct {
		Type        PacketType
		Source      Addr
		Destination Addr
		Data        []byte
	}
)

// PacketTypes
const (
	PacketTypeClose PacketType = iota + 1
	PacketTypeData
	PacketTypeOpen
)
