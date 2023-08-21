package parser

import (
	"fmt"
	"strings"
)

func (p *parser) procedure(procedure string) {
	tokens := strings.Fields(procedure)

	procedureName := tokens[0]
	fmt.Println(procedureName)
}
