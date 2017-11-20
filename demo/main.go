package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	worker1 := js.Global.Get("Worker").New("/github.com/golang-book/distributed-systems/demo/worker1/worker1.js?id=worker1")
	worker1.Set("onmessage", func(evt *js.Object) {
		str := ArrayBufferToString(evt.Get("data"))
		js.Global.Get("console").Call("log", "[main]", str)
	})
	worker1.Call("postMessage", js.S{"open", "sid1", nil})
	worker1.Call("postMessage", js.S{"send", "sid1", []byte("Hello World\n")})

	worker2 := js.Global.Get("Worker").New("/github.com/golang-book/distributed-systems/demo/worker1/worker1.js?id=worker2")
	worker2.Set("onmessage", func(evt *js.Object) {
		str := ArrayBufferToString(evt.Get("data"))
		js.Global.Get("console").Call("log", "[main]", str)
	})
	worker2.Call("postMessage", js.S{"open", "sid1", nil})
	worker2.Call("postMessage", js.S{"send", "sid1", []byte("Hello World\n")})
}

// ArrayBufferToBytes converts a javascript array buffer into a slice of bytes
func ArrayBufferToBytes(object *js.Object) []byte {
	return js.Global.Get("Uint8Array").New(object).Interface().([]byte)
}

// ArrayBufferToString converts a javascript array buffer into a string
func ArrayBufferToString(object *js.Object) string {
	return string(js.Global.Get("Uint8Array").New(object).Interface().([]byte))
}
