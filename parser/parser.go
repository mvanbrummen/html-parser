package parser

import "gitlab.com/mvanbrummen/browser-engine/dom"

type Parser interface {
	NextChar() rune
	StartsWith(str string) bool
	EOF() bool
	ConsumeChar() rune
	ConsumeWhile(predicate func(r rune) bool) string
	ConsumeWhitespace()
	ParseTagName() string
	ParseNode() dom.Node
	ParseText() dom.Node
	ParseElement() dom.Node
	ParseAttr() (string, string)
	ParseAttrValue() string
	ParseAttributes() dom.AttrMap
	ParseNodes() []*dom.Node

	Parse(source string) *dom.Node
}

type DOMParser struct {
	pos    uint
	source string
}

func NewDOMParser(source string) *DOMParser {
	return &DOMParser{
		0,
		source,
	}
}

func (p *DOMParser) NextChar() rune {
	return []rune(p.source)[p.pos+1]
}
