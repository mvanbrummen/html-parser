package parser

import "testing"

func TestParserNextChar(t *testing.T) {
	p := NewDOMParser("abc")

	result := p.NextChar()

	if result != 'b' {
		t.Errorf("Expected %c but got %c.", 'b', result)
	}
}
