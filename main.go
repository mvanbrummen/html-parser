package main

import (
	"gitlab.com/mvanbrummen/browser-engine/dom"
	"gitlab.com/mvanbrummen/browser-engine/parser"
)

func main() {
	html := `
	<html>
		<body>
			<h1>Title</h1>
			<div id="main" class="test">
				<p>Hello <em>world</em>!</p>
			</div>
		</body>
	</html>
	`
	tree := parser.Parse(html)

	dom.PrintTree(tree)
}
