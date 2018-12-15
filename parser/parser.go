package parser

import (
	"fmt"
	"strings"

	"gitlab.com/mvanbrummen/browser-engine/dom"
)

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
	if p.EOF() {
		panic(fmt.Sprintf("Cannot get %d for %s end of file", p.pos, p.source))
	}
	return []rune(p.source)[p.pos+1]
}

func (p *DOMParser) EOF() bool {
	return p.pos == uint(len(p.source))-1
}

func (p *DOMParser) StartsWith(str string) bool {
	return strings.HasPrefix(p.source, str)
}

func (p *DOMParser) ConsumeChar() rune {
	if p.EOF() {
		panic(fmt.Sprintf("Cannot get %d for %s end of file", p.pos, p.source))
	}
	p.pos++
	return []rune(p.source)[p.pos]
}
