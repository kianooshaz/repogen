package parser

import (
	"fmt"
	"strings"
)

// srcPrefix is the prefix used to identify source comments.
var srcPrefix = "src:"

// comment responsible for parsing comments.
func (p *parser) comment() error {
	r, eof := p.nextRune()
	if eof {
		return nil
	}

	if r != minus {
		return fmt.Errorf("line %d of source file is undefind!\n", p.lineNumber)
	}

	comment, _ := p.nextLine()
	comment = strings.TrimSpace(comment)

	// Check if the comment starts with the source prefix.
	if strings.HasPrefix(comment, srcPrefix) {
		// Store the parsed source in the parser's lastSrc field, removing any leading/trailing spaces.
		p.lastSrc = strings.TrimSpace(comment[len(srcPrefix):])
	}

	return nil
}
