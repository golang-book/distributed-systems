package components

import "github.com/gopherjs/vecty"
import "github.com/gopherjs/vecty/elem"

type FunctionComponent struct {
	vecty.Core
}

func (f *FunctionComponent) Render() *vecty.HTML {
	return elem.Div(
		elem.Div(vecty.Attribute("class", "function-name"),
			vecty.Text("factorial"),
		),
		elem.OrderedList(vecty.Attribute("class", "function-parameters"),
			elem.ListItem(vecty.Attribute("class", "parameter"),
				elem.Span(vecty.Attribute("class", "parameter-name"),
					vecty.Text("x"),
				),
				elem.Span(vecty.Attribute("class", "parameter-type"),
					vecty.Text("int"),
				),
			),
		),
		elem.OrderedList(vecty.Attribute("class", "function-results"),
			elem.ListItem(vecty.Attribute("class", "parameter"),
				elem.Span(vecty.Attribute("class", "parameter-type"),
					vecty.Text("int"),
				),
			),
		),
	)
}
