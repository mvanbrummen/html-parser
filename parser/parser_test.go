package parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mvanbrummen/html-parser/dom"
)

func TestDOMParser_NextChar(t *testing.T) {
	tests := []struct {
		pos         uint
		str         string
		expected    rune
		shouldPanic bool
	}{
		{0, "abc", 'a', false},
		{1, "abc", 'b', false},
		{2, "abc", 'c', false},
		{3, "abc", -1, true},
	}

	for _, test := range tests {
		p := &DOMParser{test.pos, test.str}

		if test.shouldPanic {
			assert.Panics(t, func() { p.NextChar() }, "Expected to panic")
		} else {
			result := p.NextChar()

			assert.Equal(t, result, test.expected, "Should be equal")
		}
	}
}

func TestDOMParser_EOF(t *testing.T) {
	tests := []struct {
		pos   uint
		str   string
		isEOF bool
	}{
		{3, "abc", true},
		{2, "abc", false},
		{1, "abc", false},
		{0, "abc", false},
		{1, "a", true},
	}

	for _, test := range tests {
		p := &DOMParser{test.pos, test.str}

		result := p.EOF()

		assert.Equal(t, result, test.isEOF, "Expected %v but got %v.", test.isEOF, result)
	}
}

func TestDOMParser_StartsWith(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Should return true when starts with", fields{0, "feather"}, args{"feath"}, true},
		{"Should return true when starts with later pos", fields{3, "feather"}, args{"ther"}, true},
		{"Should return true when starts with whole word", fields{0, "feather"}, args{"feather"}, true},
		{"Should return false when does not start with", fields{0, "feather"}, args{"moo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			if got := p.StartsWith(tt.args.str); got != tt.want {
				t.Errorf("DOMParser.StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDOMParser_ConsumeChar(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name        string
		fields      fields
		want        rune
		shouldPanic bool
	}{
		{"Should return rune when consume char", fields{0, "abc"}, 'a', false},
		{"Should return rune when consume char", fields{1, "abc"}, 'b', false},
		{"Should return rune when consume char", fields{2, "abc"}, 'c', false},
		{"Should panic when consume char is EOF", fields{3, "abc"}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			if tt.shouldPanic {
				assert.Panics(t, func() { p.ConsumeChar() }, "Should panic")
			} else {
				result := p.ConsumeChar()
				assert.Equal(t, result, tt.want, "DOMParser.ConsumeChar() = %v, want %v", result, tt.want)
				assert.Equal(t, p.pos, tt.fields.pos+1, "Pos should have advanced by 1")
			}
		})
	}
}

func TestDOMParser_ConsumeWhile(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	type args struct {
		predicate func(rune) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Should consume while predicate is true", fields{0, "aaabc"}, args{func(r rune) bool { return r == 'a' }}, "aaa"},
		{"Should consume while predicate is true and EOF is reached", fields{3, "aaabb"}, args{func(r rune) bool { return r == 'b' }}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			if got := p.ConsumeWhile(tt.args.predicate); got != tt.want {
				t.Errorf("DOMParser.ConsumeWhile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDOMParser_ConsumeWhitespace(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name    string
		fields  fields
		wantPos uint
	}{
		{"Should consume whitespace", fields{0, "   peacock"}, 3},
		{"Should consume nothing when no whitespace", fields{0, "peacock"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			p.ConsumeWhitespace()

			assert.Equal(t, p.pos, tt.wantPos, "Should be equal")
		})
	}
}

func TestDOMParser_ParseTagName(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Should parse tag name", fields{0, "div>"}, "div"},
		{"Should not parse tag name with special chars", fields{0, "&^%$>"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			if got := p.ParseTagName(); got != tt.want {
				t.Errorf("DOMParser.ParseTagName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDOMParser_ParseText(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   *dom.Node
	}{
		{"Should parse text node", fields{0, "hello<em>"}, dom.NewTextNode("hello")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			if got := p.ParseText(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DOMParser.ParseText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDOMParser_ParseElement(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"Should parse element node",
			fields{0, `<div id="test"><p>hello</p></div>`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}

			assert.NotNil(t, p.ParseElement())
		})
	}
}

func TestDOMParser_ParseAttrValue(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Should parse attribute value", fields{0, `"value"`}, "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			if got := p.ParseAttrValue(); got != tt.want {
				t.Errorf("DOMParser.ParseAttrValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDOMParser_ParseAttr(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{"Should parse attribute", fields{0, `id="testId"`}, "id", "testId"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			got, got1 := p.ParseAttr()
			if got != tt.want {
				t.Errorf("DOMParser.ParseAttr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DOMParser.ParseAttr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDOMParser_ParseAttributes(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   dom.AttrMap
	}{
		{"Should parse attributes", fields{0, ` name="foo" style="red">`}, dom.AttrMap{"name": "foo", "style": "red"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			if got := p.ParseAttributes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DOMParser.ParseAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDOMParser_ParseNode(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Should parse text node", fields{0, "hello<em>"}},
		{"Should parse element node", fields{0, "<div></div>"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			assert.NotNil(t, p.ParseNode())
		})
	}
}

func TestDOMParser_ParseNodes(t *testing.T) {
	type fields struct {
		pos    uint
		source string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Should parse nodes", fields{0, "<div><p>hi</p></div>"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DOMParser{
				pos:    tt.fields.pos,
				source: tt.fields.source,
			}
			assert.NotEmpty(t, p.ParseNodes())
		})
	}
}

func Test_Parse(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Should parse successfully with root node", args{`<html language="en"><div id="test1"><p>Hi <strong>you</strong></p></div></html>`}},
		{"Should parse successfully without root node", args{`<div language="en"><div id="test1"><p>Hi <strong>you</strong></p></div></div><div></div>`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.NotNil(t, Parse(tt.args.source))
		})
	}
}
