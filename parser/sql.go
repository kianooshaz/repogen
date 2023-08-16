package parser

import (
	"fmt"
	"strings"
	"unicode"
)

// Constants for prefixes and suffix used in SQL parsing.
var (
	createTablePrefix     = "CREATE TABLE"
	createProcedurePrefix = "CREATE PROCEDURE"
	createProcedureSuffix = "END"
)

// sql responsible for parsing SQL statements.
func (p *parser) sql(r rune) error {
	var eof bool
	var query string

	line := p.line

	// Iterate through the input runes.
	for {
		query += strings.ToLower(string(r))

		// Check if the query matches the CREATE TABLE prefix.
		if strings.ToUpper(query) == createTablePrefix {
			return p.createTable()
		}

		// Check if the query matches the CREATE PROCEDURE prefix.
		if strings.ToUpper(query) == createProcedurePrefix {
			return p.createProcedure()
		}

		r, eof = p.next()
		// A line has started that did not meet our expectations
		if eof {
			return fmt.Errorf("line %d of source file is undefind!\n", line)
		}
	}
}

// createTable responsible for parsing CREATE TABLE statements.
func (p *parser) createTable() error {
	var r rune
	var eof bool
	var table string

	// Iterate through the runes to find the start of the table name.
	for {
		r, eof = p.next()
		if eof {
			return nil
		}

		if unicode.IsLetter(r) {
			break
		}
	}

	// Iterate through the runes to extract the table name.
	for {
		table += string(r)

		r, eof = p.next()
		if eof {
			return nil
		}

		if unicode.IsSpace(r) || isNewLine(r) {
			break
		}
	}

	//TODO
	fmt.Println("table:", table)

	return nil
}

// createProcedure responsible for parsing CREATE PROCEDURE statements.
func (p *parser) createProcedure() error {
	var r rune
	var eof bool
	var procedure string

	// Iterate through the runes to find the start of the procedure statement.
	for {
		r, eof = p.next()
		if eof {
			return nil
		}

		if unicode.IsLetter(r) {
			break
		}
	}

	// Iterate through the runes to extract the procedure statement.
	for {
		procedure += string(r)

		// Check if the procedure name ends with the specified suffix.
		if strings.HasSuffix(strings.ToUpper(procedure), createProcedureSuffix) {
			break
		}

		r, eof = p.next()
		if eof {
			return nil
		}
	}

	// TODO
	fmt.Println("procedure: ", procedure)

	return nil
}
