package parser

import (
	"fmt"
	"io"
	"unicode"
)

// parse responsible for initiating the parsing process.
func (p *parser) parse() error {
	r, eof := p.next()
	if eof {
		return nil
	}

	// Based on the current rune, determine the parsing action to take.
	switch {
	case r == minus:
		// If the current rune is '-', invoke the comment parsing.
		if err := p.comment(); err != nil {
			return err
		}
	case unicode.IsLetter(r):
		// If the current rune is a letter, initiate SQL statement parsing.
		if err := p.sql(r); err != nil {
			return err
		}
	case !unicode.IsSpace(r):
		// If the current rune is not a space and not part of a recognized construct, return an error.
		return fmt.Errorf("line %d of source file is undefind!\n", p.line)
	}

	// Continue parsing the remaining content.
	return p.parse()
}

// next responsible for retrieving the next rune from the input buffer.
func (p *parser) next() (rune, bool) {
	r, _, err := p.buffer.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0, true
		}
		panic(err)
	}

	// Check for a new line character and update the line count accordingly.
	if isNewLine(r) {
		p.line++
	}

	return r, false
}

func (p *parser) undo(r rune) {
	if isNewLine(r) {
		p.line--
	}

	// Unread the last rune from the buffer.
	if err := p.buffer.UnreadRune(); err != nil {
		panic(err)
	}
}
