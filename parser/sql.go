package parser

import (
	"fmt"
	"strings"
)

// sql responsible for parsing SQL statements.
func (p *parser) sql(r rune) error {
	lineNumber := p.lineNumber

	if r != 'c' && r != 'C' {
		return fmt.Errorf("line %d of source file is undefind!\n", lineNumber)
	}

	query, eof := p.nextSemicolon()
	if eof {
		return fmt.Errorf("line %d of source file is undefind!\n", lineNumber)
	}

	query = string(r) + query

	words := strings.Fields(query)
	if len(words) < 2 {
		return fmt.Errorf("line %d of source file is undefind!\n", lineNumber)
	}

	switch {
	case strings.ToUpper(words[0]) == "CREATE" && strings.ToUpper(words[1]) == "TABLE":
		return p.createTable(query)
	case strings.ToUpper(words[0]) == "CREATE" && strings.ToUpper(words[1]) == "PROCEDURE":
		return p.createProcedure(query)
	default:
		return fmt.Errorf("line %d of source file is undefind!\n", lineNumber)
	}
}

// createTable responsible for parsing CREATE TABLE statements.
func (p *parser) createTable(query string) error {
	// TODO implement
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(query)
	return nil
}

// createProcedure responsible for parsing CREATE PROCEDURE statements.
func (p *parser) createProcedure(query string) error {
	// TODO implement
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(query)
	return nil
}
