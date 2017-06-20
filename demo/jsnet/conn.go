package jsnet

import "github.com/gopherjs/gopherjs/js"

type Conn struct {
	id     string
	object *js.Object
}

// // Close closes the connection
// func (c *Conn) Close() error {
// 	if c.object != nil {
// 		c.object.Call("postMessage", js.S{"close", c.id})
// 		c.object = nil
// 	}
// 	return nil
// }

// // LocalAddr returns the local address
// func (c *Conn) LocalAddr() net.Addr {
// 	return nil
// }

// // Read reads from the connection
// func (c *Conn) Read() ([]byte, error) {
// 	return nil, fmt.Errorf("not implemented")
// }

// var _ net.Conn = (*Conn)(nil)

// func Dial(name string) (*Conn, error) {
// 	return nil, fmt.Errorf("not implemented")
// }
