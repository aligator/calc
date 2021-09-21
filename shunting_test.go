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
				{Type: calc.Number, Value: "1"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "2"},
				{Type: calc.Operator, Value: "-"},
				{Type: calc.Number, Value: "3"},
				{Type: calc.Operator, Value: "-"},
				{Type: calc.Constant, Value: "33"},
			},
			want: calc.Stack{
				{Type: calc.Number, Value: "1"},
				{Type: calc.Number, Value: "2"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "3"},
				{Type: calc.Operator, Value: "-"},
				{Type: calc.Constant, Value: "33"},
				{Type: calc.Operator, Value: "-"},
			},
		},
		{
			name: "with several matching Parentheses",
			input: calc.Stack{ // (( 1 + 2 * (77 + 55)) + 3)
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Number, Value: "1"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "2"},
				{Type: calc.Operator, Value: "*"},
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Number, Value: "77"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "55"},
				{Type: calc.Rparen, Value: ")"},
				{Type: calc.Rparen, Value: ")"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "3"},
				{Type: calc.Rparen, Value: ")"},
			},
			want: calc.Stack{
				{Type: calc.Number, Value: "1"},
				{Type: calc.Number, Value: "2"},
				{Type: calc.Number, Value: "77"},
				{Type: calc.Number, Value: "55"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Operator, Value: "*"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "3"},
				{Type: calc.Operator, Value: "+"},
			},
		},
		{
			name: "with some not matching Parentheses",
			input: calc.Stack{ // (( 1 + 2 * ((77 + 55) + 3))
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Number, Value: "1"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "2"},
				{Type: calc.Operator, Value: "*"},
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Number, Value: "77"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "55"},
				{Type: calc.Rparen, Value: ")"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "3"},
				{Type: calc.Rparen, Value: ")"},
				{Type: calc.Rparen, Value: ")"},
			},
			wantErr: true,
		},
		{
			name: "with other not matching Parentheses",
			input: calc.Stack{ // ( 1 + 2 * ((77 + 55)) + 3))
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Number, Value: "1"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "2"},
				{Type: calc.Operator, Value: "*"},
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Lparen, Value: "("},
				{Type: calc.Number, Value: "77"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "55"},
				{Type: calc.Rparen, Value: ")"},
				{Type: calc.Rparen, Value: ")"},
				{Type: calc.Operator, Value: "+"},
				{Type: calc.Number, Value: "3"},
				{Type: calc.Rparen, Value: ")"},
				{Type: calc.Rparen, Value: ")"},
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
