package jsnet

import (
	"errors"
	"fmt"
	"net"
	"time"
)

const newConnectionTimeout = time.Second * 10

type conn struct {
	local, remote Addr
}

func newConn(local, remote Addr) *conn {
	return &conn{
		local:  local,
		remote: remote,
	}
}

func (c *conn) send(packetType PacketType, data []byte) error {
	return router.Send(Packet{
		Type:        packetType,
		Source:      c.local,
		Destination: c.remote,
		Data:        data,
	})
}

func (c *conn) Close() error {
	return c.send(PacketTypeClose, nil)
}

func (c *conn) Handle(packet Packet) error {
	return errors.New("not implemented")
}

func (c *conn) LocalAddr() net.Addr {
	return c.local
}

func (c *conn) Read(p []byte) (int, error) {
	return 0, errors.New("not implemented")
}

func (c *conn) RemoteAddr() net.Addr {
	return c.remote
}

func (c *conn) SetDeadline(deadline time.Time) error {
	panic("not implemented")
}

func (c *conn) SetReadDeadline(deadline time.Time) error {
	panic("not implemented")
}

func (c *conn) SetWriteDeadline(deadline time.Time) error {
	panic("not implemented")
}

func (c *conn) Write(p []byte) (int, error) {
	err := c.send(PacketTypeData, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

type listener struct {
	addr Addr

	incoming    chan *conn
	connections map[Addr]*conn
}

func newListener(addr Addr) *listener {
	return &listener{
		addr:        addr,
		connections: make(map[Addr]*conn),
		incoming:    make(chan *conn, 1),
	}
}

func (l *listener) Accept() (net.Conn, error) {
	return <-l.incoming, nil
}

func (l *listener) Addr() net.Addr {
	return l.addr
}

func (l *listener) Close() error {
	return router.Unregister(l.addr)
}

func (l *listener) Handle(packet Packet) error {
	switch packet.Type {
	case PacketTypeClose:
		c, ok := l.connections[packet.Source]
		if ok {
			delete(l.connections, packet.Source)
		}
		return c.Close()
	case PacketTypeData:
		c, ok := l.connections[packet.Source]
		if !ok {
			return fmt.Errorf("no connection source=%s destination=%s", packet.Source, packet.Destination)
		}
		return c.Handle(packet)
	case PacketTypeOpen:
		c, ok := l.connections[packet.Source]
		if !ok {
			c = newConn(l.addr, packet.Source)
			l.connections[packet.Source] = c

			select {
			case l.incoming <- c:
			case <-time.After(newConnectionTimeout):
				return fmt.Errorf("listener timed out accepting new connections")
			}
		}
		return nil
	default:
		panic(fmt.Sprintf("unknown packet type: %d", packet.Type))
	}
}

// Listen creates a listener on the given port
func Listen(port uint16) (net.Listener, error) {
	l := newListener(Addr{
		Host: router.host,
		Port: port,
	})
	return l, router.Register(l.addr, l)
}

// type listenerCloser struct {
// 	listener *Listener
// 	conn     *Conn
// }

// func (c listenerCloser) Close() error {
// 	c.listener.mu.Lock()
// 	delete(c.listener.connections, c.conn.destination)
// 	c.listener.mu.Unlock()
// 	close(c.conn.outgoing)
// 	return nil
// }

// // A Listener accepts new connections
// type Listener struct {
// 	id       string
// 	incoming chan string

// 	mu          sync.Mutex
// 	connections map[string]*Conn
// }

// var _ net.Listener = (*Listener)(nil)

// // Accept accepts a new connection
// func (li *Listener) Accept() (net.Conn, error) {
// 	id := <-li.incoming
// 	cin := make(chan []byte)
// 	cout := make(chan []byte)
// 	go func() {
// 		for data := range cout {
// 			buf := js.NewArrayBuffer(data)
// 			js.Global.Call("postMessage", buf, []interface{}{buf})
// 		}
// 	}()

// 	c := NewConn(id, li.id, cin, cout)
// 	c.closer = listenerCloser{
// 		listener: li,
// 		conn:     c,
// 	}
// 	li.mu.Lock()
// 	li.connections[id] = c
// 	li.mu.Unlock()
// 	return c, nil
// }

// // Addr returns the address
// func (li *Listener) Addr() net.Addr {
// 	return Addr{li.id}
// }

// // Close closes the listener
// func (li *Listener) Close() error {
// 	li.mu.Lock()
// 	for _, conn := range li.connections {
// 		conn.Close()
// 	}
// 	li.connections = nil
// 	li.mu.Unlock()
// 	return nil
// }

// func (li *Listener) dispatch(cmd string, id string, data *js.Object) {
// 	switch cmd {
// 	case "open":
// 		li.incoming <- id
// 	case "close":
// 		li.mu.Lock()
// 		if c, ok := li.connections[id]; ok {
// 			c.Close()
// 			delete(li.connections, id)
// 		}
// 		li.mu.Unlock()
// 	case "send":
// 		li.mu.Lock()
// 		c, ok := li.connections[id]
// 		li.mu.Unlock()
// 		if ok {
// 			c.incoming <- data.Interface().([]byte)
// 		}
// 	default:
// 		log.Println("["+li.id+"] unknown command", cmd, data)
// 	}
// }

// // Listen creates a new Listener
// func Listen() (*Listener, error) {
// 	u, err := url.Parse(js.Global.Get("location").Get("href").String())
// 	if err != nil {
// 		return nil, err
// 	}

// 	li := &Listener{
// 		id:          u.Query().Get("id"),
// 		incoming:    make(chan string),
// 		connections: make(map[string]*Conn),
// 	}
// 	js.Global.Set("onmessage", func(evt *js.Object) {
// 		cmd := evt.Get("data").Index(0).String()
// 		sid := evt.Get("data").Index(1).String()
// 		data := evt.Get("data").Index(2)
// 		go li.dispatch(cmd, sid, data)
// 	})
// 	return li, nil
// }
