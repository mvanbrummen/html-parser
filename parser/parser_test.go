package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserNextChar(t *testing.T) {
	tables := []struct {
		pos         uint
		str         string
		expected    rune
		shouldPanic bool
	}{
		{0, "abc", 'b', false},
		{1, "abc", 'c', false},
		{2, "abc", -1, true},
	}

	for _, table := range tables {
		p := &DOMParser{table.pos, table.str}

		if table.shouldPanic {
			assert.Panics(t, func() { p.NextChar() }, "Expected to panic")
		} else {
			result := p.NextChar()

			assert.Equal(t, result, table.expected, "Should be equal")
		}
	}
}

func TestParserEOF(t *testing.T) {
	tables := []struct {
		pos   uint
		str   string
		isEOF bool
	}{
		{2, "abc", true},
		{1, "abc", false},
		{0, "abc", false},
		{0, "a", true},
	}

	for _, table := range tables {
		p := &DOMParser{table.pos, table.str}

		result := p.EOF()

		assert.Equal(t, result, table.isEOF, "Expected %v but got %v.", table.isEOF, result)
	}
}
