package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	worker1 := js.Global.Get("Worker").New("/github.com/golang-book/distributed-systems/demo/worker1/worker1.js", js.M{
		"name": "worker1",
	})
	worker1.Set("onmessage", func(evt *js.Object) {
		js.Global.Get("console").Call("log", "[main]", evt.Get("data"))
	})
	worker1.Call("postMessage", js.S{"open", "sid1", nil})
	worker1.Call("postMessage", js.S{"send", "sid1", []byte("Hello World\n")})

	worker2 := js.Global.Get("Worker").New("/github.com/golang-book/distributed-systems/demo/worker1/worker1.js", js.M{
		"name": "worker2",
	})
	worker2.Set("onmessage", func(evt *js.Object) {
		js.Global.Get("console").Call("log", "[main]", evt.Get("data"))
	})
	worker2.Call("postMessage", js.S{"open", "sid1", nil})
	worker2.Call("postMessage", js.S{"send", "sid1", []byte("Hello World\n")})
}
