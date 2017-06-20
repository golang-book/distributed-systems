package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Workspace struct {
	vecty.Core
}

func (w *Workspace) Render() *vecty.HTML {
	return elem.Div(vecty.Attribute("class", "workspace"),
		elem.Style(vecty.Text(`
* {
	font-family: 'Go Mono';
	font-weight: normal;
	font-style: normal;
	font-size: 16px;
}
.code-block {
	transform: scale(0.8);
}
.view-line.read-only {
	color: #999;
}
`)),
		new(FunctionComponent),
	)
}
