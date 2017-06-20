package jsnet

import (
	"log"
	"net"
	"sync"
	"time"

	"io"

	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

type serverConn struct {
	id       string
	incoming chan []byte

	ready []byte
}

func (c *serverConn) Close() error {
	js.Global.Call("postMessage", js.S{"close", c.id, nil})
	return nil
}

func (c *serverConn) LocalAddr() net.Addr {
	return Addr{c.id}
}

func (c *serverConn) RemoteAddr() net.Addr {
	return Addr{c.id}
}

func (c *serverConn) Read(data []byte) (int, error) {
	if len(c.ready) == 0 {
		var ok bool
		c.ready, ok = <-c.incoming
		if !ok {
			return 0, io.EOF
		}
	}

	if len(data) > len(c.ready) {
		copy(data, c.ready)
		c.ready = c.ready[len(data):]
		return len(data), nil
	}

	copy(data, c.ready)
	n := len(c.ready)
	c.ready = nil
	return n, nil
}

func (c *serverConn) SetDeadline(deadline time.Time) error {
	return fmt.Errorf("not implemented")
}

func (c *serverConn) SetReadDeadline(deadline time.Time) error {
	return fmt.Errorf("not implemented")
}

func (c *serverConn) SetWriteDeadline(deadline time.Time) error {
	return fmt.Errorf("not implemented")
}

func (c *serverConn) Write(data []byte) (int, error) {
	buf := js.NewArrayBuffer(data)
	js.Global.Call("postMessage", js.S{"send", c.id, buf}, js.S{buf})
	return len(data), nil
}

type Listener struct {
	name     string
	incoming chan string

	mu          sync.Mutex
	connections map[string]*serverConn
}

var _ net.Listener = (*Listener)(nil)

func (li *Listener) Accept() (net.Conn, error) {
	id := <-li.incoming
	c := &serverConn{
		id:       id,
		incoming: make(chan []byte),
	}
	li.mu.Lock()
	li.connections[id] = c
	li.mu.Unlock()
	return c, nil
}

// Addr returns the address
func (li *Listener) Addr() net.Addr {
	return Addr{li.name}
}

// Close closes the listener
func (li *Listener) Close() error {
	for _, conn := range li.connections {
		conn.Close()
	}
	return nil
}

func (li *Listener) dispatch(cmd string, id string, data *js.Object) {
	switch cmd {
	case "open":
		li.incoming <- id
	case "close":
		li.mu.Lock()
		if c, ok := li.connections[id]; ok {
			c.Close()
			delete(li.connections, id)
		}
		li.mu.Unlock()
	case "send":
		li.mu.Lock()
		c, ok := li.connections[id]
		li.mu.Unlock()
		if ok {
			c.incoming <- data.Interface().([]byte)
		}
	default:
		log.Println("["+li.name+"] unknown command", cmd, data)
	}
}

func Listen() *Listener {
	li := &Listener{
		name:        js.Global.Get("name").String(),
		incoming:    make(chan string),
		connections: make(map[string]*serverConn),
	}
	js.Global.Set("onmessage", func(evt *js.Object) {
		cmd := evt.Get("data").Index(0).String()
		sid := evt.Get("data").Index(1).String()
		data := evt.Get("data").Index(2)
		go li.dispatch(cmd, sid, data)
	})
	return li
}
