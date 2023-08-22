package parser

import (
	"bufio"
	"github.com/kianooshaz/repogen/protocol"
	"strings"
)

// parser is a struct responsible for parsing the input source according to the specified protocol.
type parser struct {
	buffer     *bufio.Reader // Buffered reader to handle an input source.
	generator  protocol.Generator
	lineNumber uint   // lineNumber counter to keep track of the current line being parsed.
	lastSrc    string // A string to store the last parsed source.
}

// New creates a new instance of the parser.
func New(source string) protocol.Parser {
	return &parser{
		buffer:     bufio.NewReader(strings.NewReader(source)),
		lineNumber: 1,
	}
}
