package calc

import "testing"

func TestContainsLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsLetter(tt.args.s); got != tt.want {
				t.Errorf("ContainsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{name: "no input", input: "", wantErr: true},
		{name: "simple plus", input: "5+4", want: 9},
		{name: "simple minus", input: "5-4", want: 1},
		{name: "simple multiplication", input: "42*2", want: 84},
		{name: "simple division", input: "42/2", want: 21},
		{name: "simple exp", input: "2^3", want: 8},
		{name: "with constant", input: "2^3+PI", want: 11.141592653589793},
		{name: "with function", input: "COS(5)", want: 0.2836621854632263},
		{name: "with parentheses", input: "2*(5+3)", want: 16},
		{name: "with more parentheses", input: "((2*(5+3))+4)*(300/100)", want: 60},
		{name: "with spaces, tabs and newlines", input: "    (  \n   (2*(  \t\t\t5+    3))+4)*  \n       (300    / 100)   ", want: 60},
		{name: "invalid calculation: ends with operator", input: "((2*(5+3))+4)+", wantErr: true},
		{name: "invalid calculation: double operator", input: "((2*(5+3))++4)", wantErr: true},
		{name: "invalid calculation: wrong parentheses", input: "((2*(5+3)+4", wantErr: true},
		{name: "invalid calculation: wrong parentheses2", input: "(2*(5+3))+4)", wantErr: true},
		{name: "invalid calculation: invalid function", input: "(2*(5+3))+4*LOOL(5)", wantErr: true},
		{name: "invalid calculation: invalid constant", input: "(2*(5+3))+4*LOOL", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solve() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveFunction(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveFunction(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveFunction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveFunction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvePostfix(t *testing.T) {
	type args struct {
		tokens Stack
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePostfix(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolvePostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolvePostfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
