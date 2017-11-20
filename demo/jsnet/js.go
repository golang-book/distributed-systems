package jsnet

import (
	"errors"
	"net/url"

	"sync"

	"github.com/gopherjs/gopherjs/js"
)

type serialSocket struct {
	sync.Mutex
	socket Socket
}

// A JSRouter implements packet routing for a web worker
type JSRouter struct {
	host   Host
	global *js.Object

	mu      sync.RWMutex
	sockets map[Addr]*serialSocket
}

// NewJSRouter creates a new JSRouter
func NewJSRouter(host Host, global *js.Object) *JSRouter {
	s := &JSRouter{
		host:    host,
		global:  global,
		sockets: make(map[Addr]*serialSocket),
	}

	global.Set("onmessage", func(evt *js.Object) {
		data := evt.Get("data")
		packet := Packet{
			Source:      AddrFromString(data.Index(0).String()),
			Destination: AddrFromString(data.Index(1).String()),
			Data:        data.Index(2).Interface().([]byte),
		}
		go s.Recv(packet)
	})

	return s
}

// Recv receives a packet
func (s *JSRouter) Recv(packet Packet) error {
	s.mu.RLock()
	socket, ok := s.sockets[packet.Destination]
	s.mu.RUnlock()
	if ok {
		socket.Lock()
		defer socket.Unlock()
		return socket.socket.Handle(packet)
	}
	return errors.New("socket not found")
}

// Register registers a new socket on port
func (s *JSRouter) Register(addr Addr, socket Socket) error {
	s.mu.Lock()
	cur, ok := s.sockets[addr]
	if ok && cur.socket == socket {
		ok = false
	}
	if !ok {
		s.sockets[addr] = &serialSocket{
			socket: socket,
		}
	}
	s.mu.Unlock()
	if ok {
		return errors.New("port taken")
	}
	return nil
}

// Unregister unregisters a socket from a port
func (s *JSRouter) Unregister(addr Addr) error {
	s.mu.Lock()
	delete(s.sockets, addr)
	s.mu.Unlock()
	return nil
}

// Send sends a packet
func (s *JSRouter) Send(packet Packet) error {
	s.mu.RLock()
	socket, ok := s.sockets[packet.Destination]
	s.mu.RUnlock()
	if ok {
		socket.Lock()
		defer socket.Unlock()
		return socket.socket.Handle(packet)
	}

	s.global.Call("postMessage", []interface{}{
		packet.Source.String(),
		packet.Destination.String(),
		packet.Data,
	})

	return nil
}

var router *JSRouter

func init() {
	u, err := url.Parse(js.Global.Get("location").Get("href").String())
	if err != nil {
		panic(err)
	}
	host := HostFromString(u.Query().Get("host"))
	router = NewJSRouter(host, js.Global)
}
