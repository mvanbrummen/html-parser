package main

import (
	"fmt"
	"strings"
)

// A AttrMap is a html attribute map type
type AttrMap = map[string]string

// A Node represents a html node
type Node struct {
	Children []*Node
	NodeType interface{}
}

// An Element represents a html element node type
type Element struct {
	TagName    string
	Attributes AttrMap
}

func printTree(node *Node) {
	printTreeWithIndentation(node, 0)
}

func printTreeWithIndentation(node *Node, indentation int) {
	switch n := node.NodeType.(type) {
	case string:
		fmt.Printf("%stext: %s\n", strings.Repeat(" ", indentation), n)
	case Element:
		fmt.Printf("%stag: %s", strings.Repeat(" ", indentation), n.TagName)

		if n.Attributes != nil {
			fmt.Printf(" attrs: %v\n", n.Attributes)
		} else {
			fmt.Println()
		}

		indentation++

		for _, c := range node.Children {
			printTreeWithIndentation(c, indentation)
		}
	default:
		panic("Unknown type!")
	}
}

func text(str string) *Node {
	return &Node{
		[]*Node{},
		str,
	}
}

func element(tagName string, attributes AttrMap, children []*Node) *Node {
	return &Node{
		children,
		Element{
			tagName,
			attributes,
		},
	}
}

func main() {
	tree := element("html", map[string]string{
		"language": "en",
	}, []*Node{element("div", nil, []*Node{element("div", nil, []*Node{text("Hello"), text("Some Text"), text("Some more text")})})})

	printTree(tree)
}
