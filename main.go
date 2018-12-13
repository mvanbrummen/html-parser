package main

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

}
