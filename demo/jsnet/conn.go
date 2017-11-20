package jsnet

import (
	"io"
	"net"
	"sync"
	"time"
)

type timeoutError struct{}

func (t *timeoutError) Error() string {
	return "timeout"
}
func (t *timeoutError) Timeout() bool {
	return true
}
func (t *timeoutError) Temporary() bool {
	return false
}

type Conn struct {
	source, destination Addr
	incoming            chan []byte
	outgoing            chan []byte
	closer              io.Closer

	mu                          sync.Mutex
	readDeadline, writeDeadline time.Time
	ready                       []byte
}

// NewConn creates a new connection
func NewConn(
	source, destination Addr,
	incoming chan []byte,
	outgoing chan []byte,
) *Conn {
	return &Conn{
		source:      source,
		destination: destination,
		incoming:    incoming,
		outgoing:    outgoing,
	}
}

// Close closes the connection
func (c *Conn) Close() error {
	if c.closer == nil {
		return nil
	}
	return c.closer.Close()
}

// LocalAddr returns the local address
func (c *Conn) LocalAddr() net.Addr {
	return c.source
}

// RemoteAddr returns the remote address
func (c *Conn) RemoteAddr() net.Addr {
	return c.destination
}

func (c *Conn) Read(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	return 0, nil

	// c.mu.Lock()
	// readDeadline := c.readDeadline
	// c.mu.Unlock()

	// finished := 0

	// c.mu.Lock()
	// switch {
	// case len(c.ready) > len(data):
	// 	copy(data, c.ready)
	// 	c.ready = c.ready[len(data):]
	// 	finished = len(data)
	// case len(c.ready) > 0:
	// 	copy(data, c.ready)
	// 	c.ready = nil
	// 	finished = len(c.ready)
	// }
	// c.mu.Unlock()

	// for finished == 0 {
	// 	if readDeadline.IsZero() {
	// 		read := <-c.incoming
	// 		copy(data, read)

	// 	}
	// }

	// return finished, nil

	// if finished > 0 {
	// 	return finished, nil
	// }

	// if readDeadline.IsZero() {
	// 	read = <-c.incoming
	// 	c.outgoing <- data
	// 	return len(data), nil
	// }

	// if len(ready) > len(data) {

	// }
	// c.mu.Unlock()

	// if len(c.ready) == 0 {
	// 	var ok bool
	// 	c.ready, ok = <-c.incoming
	// 	if !ok {
	// 		return 0, io.EOF
	// 	}
	// }

	// if len(data) > len(c.ready) {
	// 	copy(data, c.ready)
	// 	c.ready = c.ready[len(data):]
	// 	return len(data), nil
	// }

	// copy(data, c.ready)
	// n := len(c.ready)
	// c.ready = nil
	// return n, nil
}

// SetDeadline sets the write and read deadline
func (c *Conn) SetDeadline(deadline time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.readDeadline = deadline
	c.writeDeadline = deadline
	return nil
}

// SetReadDeadline sets the read deadline
func (c *Conn) SetReadDeadline(deadline time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.readDeadline = deadline
	return nil
}

// SetWriteDeadline sets the write deadline
func (c *Conn) SetWriteDeadline(deadline time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.writeDeadline = deadline
	return nil
}

func (c *Conn) Write(data []byte) (int, error) {
	c.mu.Lock()
	writeDeadline := c.writeDeadline
	c.mu.Unlock()

	if writeDeadline.IsZero() {
		c.outgoing <- data
		return len(data), nil
	}

	now := time.Now()
	if now.After(writeDeadline) {
		return 0, new(timeoutError)
	}

	timer := time.NewTimer(writeDeadline.Sub(now))
	defer timer.Stop()

	select {
	case <-timer.C:
		return 0, new(timeoutError)
	case c.outgoing <- data:
		return len(data), nil
	}
}
