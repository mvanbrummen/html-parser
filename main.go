package main

import "gitlab.com/mvanbrummen/browser-engine/dom"

func main() {
	tree := dom.NewElementNode("html", map[string]string{
		"language": "en",
	}, []*dom.Node{
		dom.NewElementNode("div", nil, []*dom.Node{
			dom.NewElementNode("div", nil, []*dom.Node{
				dom.NewTextNode("Hello"),
				dom.NewTextNode("Some Text"),
				dom.NewTextNode("Some more text"),
			}),
		}),
	})

	dom.PrintTree(tree)
}
