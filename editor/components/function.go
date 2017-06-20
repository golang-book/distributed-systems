package components

import "github.com/gopherjs/vecty"
import "github.com/gopherjs/vecty/elem"

type FunctionComponent struct {
	vecty.Core
}

func (f *FunctionComponent) Render() *vecty.HTML {
	return elem.Div(vecty.Attribute("class", "code-block"),
		elem.Div(vecty.Attribute("class", "view-line read-only"),
			vecty.Text("func factorial(x int) int {")),

		elem.Div(vecty.Attribute("class", "view-line")),
		elem.Div(vecty.Attribute("class", "view-line read-only"),
			vecty.Text("}")),
	)
}
