package browser

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

// PrintTree pretty prints a tree
func PrintTree(node *Node) {
	printTreeWithIndentation(node, 0)
}

func printTreeWithIndentation(node *Node, indentation int) {
	repeatSpaces := strings.Repeat(" ", indentation)

	switch n := node.NodeType.(type) {
	case string:
		fmt.Printf("%stext: %s\n", repeatSpaces, n)
	case Element:
		fmt.Printf("%stag: %s", repeatSpaces, n.TagName)

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

// NewTextNode returns a new text node
func NewTextNode(str string) *Node {
	return &Node{
		[]*Node{},
		str,
	}
}

// NewElementNode returns a new Element node
func NewElementNode(tagName string, attributes AttrMap, children []*Node) *Node {
	return &Node{
		children,
		Element{
			tagName,
			attributes,
		},
	}
}
