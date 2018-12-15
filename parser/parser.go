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
	ParseNode() *dom.Node
	ParseText() *dom.Node
	ParseElement() *dom.Node
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
	return strings.HasPrefix(p.source[p.pos:], str)
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

func (p *DOMParser) ParseTagName() string {
	isAlphaNumeric := func(r rune) bool {
		return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
	}
	return p.ConsumeWhile(isAlphaNumeric)
}

func (p *DOMParser) ParseNode() *dom.Node {
	switch p.NextChar() {
	case '<':
		return p.ParseElement()
	default:
		return p.ParseText()
	}
}

func (p *DOMParser) ParseText() *dom.Node {
	isNotOpeningBracket := func(r rune) bool {
		return r != '<'
	}
	str := p.ConsumeWhile(isNotOpeningBracket)

	return dom.NewTextNode(str)
}

func assertConsumeChar(p *DOMParser, r rune) {
	char := p.ConsumeChar()
	if char != r {
		panic(fmt.Sprintf("Expected a '%c' but got a '%c' at pos %d.", r, char, p.pos))
	}
}

func (p *DOMParser) ParseElement() *dom.Node {
	assertConsumeChar(p, '<')
	tagName := p.ParseTagName()
	attributes := p.ParseAttributes()

	assertConsumeChar(p, '>')

	children := p.ParseNodes()

	assertConsumeChar(p, '<')
	assertConsumeChar(p, '/')
	if p.ParseTagName() != tagName {
		panic("Closing tagname was not " + tagName)
	}
	assertConsumeChar(p, '>')

	return dom.NewElementNode(tagName, attributes, children)
}

func (p *DOMParser) ParseNodes() []*dom.Node {
	nodes := make([]*dom.Node, 0)

	for {
		p.ConsumeWhitespace()

		if p.EOF() || p.StartsWith("</") {
			break
		}

		nodes = append(nodes, p.ParseNode())
	}

	return nodes
}

func (p *DOMParser) ParseAttributes() dom.AttrMap {
	attributes := make(dom.AttrMap, 0)

	for {
		p.ConsumeWhitespace()

		if p.NextChar() == '>' {
			break
		}
		name, value := p.ParseAttr()

		attributes[name] = value
	}

	return attributes
}

func (p *DOMParser) ParseAttrValue() string {
	openQuote := p.ConsumeChar()

	if openQuote != '"' && openQuote != '\'' {
		panic(fmt.Sprintf("Expected a quote char but got %c", openQuote))
	}

	value := p.ConsumeWhile(func(r rune) bool { return r != openQuote })

	assertConsumeChar(p, openQuote)

	return value
}

func (p *DOMParser) ParseAttr() (string, string) {
	name := p.ParseTagName()
	assertConsumeChar(p, '=')
	value := p.ParseAttrValue()

	return name, value
}

func (p *DOMParser) Parse(source string) *dom.Node {
	nodes := p.ParseNodes()

	if len(nodes) == 1 {
		return nodes[0]
	}

	return dom.NewElementNode("html", nil, nodes)
}
