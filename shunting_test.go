package calc_test

import (
	"reflect"
	"testing"

	"github.com/aligator/calc"
)

func TestShuntingYard(t *testing.T) {
	tests := []struct {
		name    string
		input   calc.Stack
		want    calc.Stack
		wantErr bool
	}{
		{
			name:  "empty input",
			input: calc.Stack{},
			want:  calc.Stack{},
		},
		{
			name: "several numbers and operators",
			input: calc.Stack{
				{Type: calc.NUMBER, Value: "1"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "2"},
				{Type: calc.OPERATOR, Value: "-"},
				{Type: calc.NUMBER, Value: "3"},
				{Type: calc.OPERATOR, Value: "-"},
				{Type: calc.CONSTANT, Value: "33"},
			},
			want: calc.Stack{
				{Type: calc.NUMBER, Value: "1"},
				{Type: calc.NUMBER, Value: "2"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "3"},
				{Type: calc.OPERATOR, Value: "-"},
				{Type: calc.CONSTANT, Value: "33"},
				{Type: calc.OPERATOR, Value: "-"},
			},
		},
		{
			name: "with several matching Parentheses",
			input: calc.Stack{ // (( 1 + 2 * (77 + 55)) + 3)
				{Type: calc.LPAREN, Value: "("},
				{Type: calc.LPAREN, Value: "("},
				{Type: calc.NUMBER, Value: "1"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "2"},
				{Type: calc.OPERATOR, Value: "*"},
				{Type: calc.LPAREN, Value: "("},
				{Type: calc.NUMBER, Value: "77"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "55"},
				{Type: calc.RPAREN, Value: ")"},
				{Type: calc.RPAREN, Value: ")"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "3"},
				{Type: calc.RPAREN, Value: ")"},
			},
			want: calc.Stack{
				{Type: calc.NUMBER, Value: "1"},
				{Type: calc.NUMBER, Value: "2"},
				{Type: calc.NUMBER, Value: "77"},
				{Type: calc.NUMBER, Value: "55"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.OPERATOR, Value: "*"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "3"},
				{Type: calc.OPERATOR, Value: "+"},
			},
		},
		{
			name: "with some not matching Parentheses",
			input: calc.Stack{ // (( 1 + 2 * ((77 + 55) + 3))
				{Type: calc.LPAREN, Value: "("},
				{Type: calc.LPAREN, Value: "("},
				{Type: calc.NUMBER, Value: "1"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "2"},
				{Type: calc.OPERATOR, Value: "*"},
				{Type: calc.LPAREN, Value: "("},
				{Type: calc.LPAREN, Value: "("},
				{Type: calc.NUMBER, Value: "77"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "55"},
				{Type: calc.RPAREN, Value: ")"},
				{Type: calc.OPERATOR, Value: "+"},
				{Type: calc.NUMBER, Value: "3"},
				{Type: calc.RPAREN, Value: ")"},
				{Type: calc.RPAREN, Value: ")"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.ShuntingYard(tt.input)

			if err != nil && !tt.wantErr {
				t.Errorf("ShuntingYard() got error %v, want nil", err)
			} else if tt.wantErr && err == nil {
				t.Error("ShuntingYard() got no error, want non-nil")
			} else {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuntingYard() = %v, want %v", got, tt.want)
			}
		})
	}
}
