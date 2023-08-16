package parser

import (
	"fmt"
	"strings"
	"unicode"
)

// srcPrefix is the prefix used to identify source comments.
var srcPrefix = "src:"

// comment responsible for parsing comments.
func (p *parser) comment() error {
	var r rune
	var eof bool
	var comment string

	r, eof = p.next()
	if eof {
		return nil
	}

	if r != minus {
		return fmt.Errorf("line %d of source file is undefind!\n", p.line)
	}

	// Iterate through runes to find the start of the comment.
	for {
		r, eof = p.next()
		if eof {
			return nil
		}

		// Once a letter is encountered, it indicates the start of the comment content.
		if unicode.IsLetter(r) {
			break
		}
	}

	// Iterate through runes to extract the comment content.
	for {
		comment += string(r)

		r, eof = p.next()
		if eof {
			return nil
		}

		// If a newline character is encountered, the comment ends.
		if isNewLine(r) {
			break
		}
	}

	// Check if the comment starts with the source prefix.
	if strings.HasPrefix(comment, srcPrefix) {
		// Store the parsed source in the parser's lastSrc field, removing any leading/trailing spaces.
		p.lastSrc = strings.TrimSpace(comment[len(srcPrefix):])
	}

	return nil
}
