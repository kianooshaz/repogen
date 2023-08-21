package parser

import (
	"bufio"
	"strings"
	"testing"
)

func Test_parser_comment(t *testing.T) {
	tests := []struct {
		name    string
		buffer  string
		wantErr bool
	}{
		{
			name:    "not exist semicolon first of buffer",
			buffer:  " src: model",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				buffer: bufio.NewReader(strings.NewReader(tt.buffer)),
			}
			if err := p.comment(); (err != nil) != tt.wantErr {
				t.Errorf("comment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
