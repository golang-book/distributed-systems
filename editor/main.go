package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Get("document").Call("write", `

<div id="container" style="width:800px;height:600px;border:1px solid grey"></div>

<script src="./assets/vs/loader.js"></script>
<script>
	require.config({ paths: { 'vs': 'assets/vs' }});
	require(['vs/editor/editor.main'], function() {
		var editor = monaco.editor.create(document.getElementById('container'), {
			fontFamily: "Go Mono",
			fontSize: 16,
			lineNumbers: false,
			value: [
				'func factorial(x int) {',
				'\treturn 0',
				'}'
			].join('\n'),
			language: 'go'
		});
	});
</script>`)
}
