package jsnet

type (
	// A Socket receives packets
	Socket interface {
		// Handle handles a packet. Handle should not be called by multiple goroutines simultaneously.
		Handle(Packet) error
	}
)
