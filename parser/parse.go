package parser

import (
	"fmt"
	"unicode"
)

// Parse responsible for initiating the parsing process.
func (p *parser) Parse() error {
	for {
		r, eof := p.nextRune()
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
			return fmt.Errorf("line %d of source file is undefind!\n", p.lineNumber)
		}
	}
}
