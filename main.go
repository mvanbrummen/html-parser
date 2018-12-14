package main

import (
	"gitlab.com/mvanbrummen/browser-engine/browser"
)

func main() {
	tree := browser.NewElementNode("html", map[string]string{
		"language": "en",
	}, []*browser.Node{
		browser.NewElementNode("div", nil, []*browser.Node{
			browser.NewElementNode("div", nil, []*browser.Node{
				browser.NewTextNode("Hello"),
				browser.NewTextNode("Some Text"),
				browser.NewTextNode("Some more text"),
			}),
		}),
	})

	browser.PrintTree(tree)
}
