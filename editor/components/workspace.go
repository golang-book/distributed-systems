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
	font-family: 'GoMonoRegular';
	font-weight: normal;
	font-style: normal;
	font-size: 16px;
}
.parameter {

}
.parameter-name {

}
.parameter-type {
	display: inline-block;
	border: 0px solid #cFc;
	border-radius: 4px;
	background: #EFE;
	padding: 4px 8px;
	box-shadow: 1px 0px #CFC;
}
.parameter-name + .parameter-type {
	margin-left: 8px;
}
`)),
		new(FunctionComponent),
	)
}
