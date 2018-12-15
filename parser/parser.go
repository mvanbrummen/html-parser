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
	ConsumeWhile(predicate func(rune) bool) string
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
	return []rune(p.source)[p.pos]
}

func (p *DOMParser) EOF() bool {
	return p.pos == uint(len(p.source))
}

func (p *DOMParser) StartsWith(str string) bool {
	return strings.HasPrefix(p.source, str)
}

func (p *DOMParser) ConsumeChar() rune {
	if p.EOF() {
		panic(fmt.Sprintf("Cannot get %d for %s end of file", p.pos, p.source))
	}
	char := []rune(p.source)[p.pos]

	p.pos++

	return char
}

func (p *DOMParser) ConsumeWhile(predicate func(rune) bool) string {
	str := ""
	for !p.EOF() && predicate(p.NextChar()) {
		str = fmt.Sprintf("%s%c", str, p.ConsumeChar())
	}

	return str
}

func (p *DOMParser) ConsumeWhitespace() {
	isWhiteSpace := func(r rune) bool { return r == ' ' }
	p.ConsumeWhile(isWhiteSpace)
}
