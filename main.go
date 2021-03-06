package main

import (
	"github.com/mvanbrummen/html-parser/dom"
	"github.com/mvanbrummen/html-parser/parser"
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
