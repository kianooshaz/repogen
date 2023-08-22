package parser

import (
	"errors"
	"github.com/kianooshaz/repogen/protocol"
	"strings"
)

func (p *parser) procedureParameters(procedure string) ([]protocol.Parameter, []protocol.Parameter, error) {
	beginIndex := strings.Index(strings.ToUpper(procedure), "BEGIN")
	if beginIndex == 0 {
		return nil, nil, errors.New("") // todo
	}
	procedure = procedure[:beginIndex]

	lParentIndex := strings.Index(procedure, "(")
	rParentIndex := strings.LastIndex(procedure, ")")

	if lParentIndex == 0 && rParentIndex == 0 {
		return nil, nil, nil
	}

	if lParentIndex == 0 || rParentIndex == 0 {
		return nil, nil, errors.New("") // todo
	}

	parametersStr := strings.Split(procedure[lParentIndex+1:rParentIndex], ",")

	var inputs []protocol.Parameter
	var outputs []protocol.Parameter
	for _, paramStr := range parametersStr {
		var input bool

		paramStr = strings.TrimSpace(paramStr)

		switch {
		case strings.HasPrefix(strings.ToUpper(paramStr), "IN"):
			input = true
			paramStr = paramStr[len("IN"):]
		case strings.HasPrefix(strings.ToUpper(paramStr), "OUT"):
			input = false
			paramStr = paramStr[len("OUT"):]
		default:
			return nil, nil, errors.New("") // todo
		}

		paramStr = strings.TrimSpace(paramStr)
		names := strings.Fields(paramStr)

		if len(names) < 2 {
			return nil, nil, errors.New("") // todo
		}

		name := names[0]
		t := paramStr[len(name):]

		if input {
			inputs = append(inputs, protocol.Parameter{
				Name: name,
				Type: t,
			})
		} else {
			outputs = append(outputs, protocol.Parameter{
				Name: name,
				Type: t,
			})
		}
	}

	return inputs, outputs, nil
}
