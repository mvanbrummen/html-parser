package main

type AttrMap = map[string]string

type Node struct {
	Children []*Node
	NodeType interface{}
}

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
