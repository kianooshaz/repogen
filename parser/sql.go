package parser

import (
	"fmt"
	"github.com/kianooshaz/repogen/protocol"
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
	if len(words) < 3 {
		return fmt.Errorf("line %d of source file is undefind!\n", lineNumber)
	}

	switch {
	case strings.ToUpper(words[0]) == "CREATE" && strings.ToUpper(words[1]) == "TABLE":
		if err := p.generator.AddEntity(words[3], p.lastSrc); err != nil {
			return fmt.Errorf("error on line %d of source file, error: %w\n", p.lineNumber, err)
		}
	case strings.ToUpper(words[0]) == "CREATE" && strings.ToUpper(words[1]) == "PROCEDURE":
		name := words[2]
		inputs, outputs, err := p.procedureParameters(query)
		if err != nil {
			return fmt.Errorf("error on line %d of source file, error: %w\n", p.lineNumber, err)
		}

		if err := p.generator.AddProcedure(protocol.Procedure{
			Name:    name,
			Inputs:  inputs,
			Outputs: outputs,
		}); err != nil {
			return fmt.Errorf("error on line %d of source file, error: %w\n", p.lineNumber, err)
		}
	default:
		return fmt.Errorf("line %d of source file is undefind!\n", lineNumber)
	}

	return nil
}
