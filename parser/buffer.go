package parser

import (
	"io"
	"strings"
)

// nextRune responsible for retrieving the nextRune rune from the input buffer.
func (p *parser) nextRune() (rune, bool) {
	r, _, err := p.buffer.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0, true
		}
		panic(err)
	}

	// Check for a new line character and update the line count accordingly.
	if isNewLine(r) {
		p.lineNumber++
	}

	return r, false
}

func (p *parser) undo(r rune) {
	if isNewLine(r) {
		p.lineNumber--
	}

	// Unread the last rune from the buffer.
	if err := p.buffer.UnreadRune(); err != nil {
		panic(err)
	}
}

func (p *parser) nextLine() (string, bool) {
	var partOfLine []byte
	var line string
	var err error

	prefix := true
	for prefix {
		partOfLine, prefix, err = p.buffer.ReadLine()
		if err != nil {
			if err == io.EOF {
				return "", true
			}
			panic(err)
		}

		line += string(partOfLine)
	}

	p.lineNumber++

	return line, false
}

func (p *parser) nextSemicolon() (string, bool) {
	text, err := p.buffer.ReadString(';')
	if err != nil {
		if err == io.EOF {
			return "", true
		}
		panic(err)
	}

	p.lineNumber += uint(strings.Count(text, "\n"))

	return text, false
}
