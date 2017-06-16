package main

import (
	"github.com/golang-book/distributed-systems/editor/components"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Main struct {
	vecty.Core
}

func (m *Main) Render() *vecty.HTML {
	return elem.Body(
		vecty.Text("Hello World"),
		new(components.Workspace),
	)
}

func main() {
	vecty.SetTitle("editor")
	vecty.AddStylesheet("https://fontlibrary.org/face/go-mono")
	vecty.RenderBody(new(Main))
}
