package calc_test

import (
	"reflect"
	"testing"

	"github.com/aligator/calc"
)

func copyStack(s calc.Stack) calc.Stack {
	var original calc.Stack
	if s != nil {
		original = make(calc.Stack, len(s))
		copy(original, s)
	}
	return original
}

func TestStack_EmptyInto(t *testing.T) {
	tests := []struct {
		name     string
		source   calc.Stack
		dest     calc.Stack
		expected calc.Stack
	}{
		{
			name: "empty one stack into an empty one",
			source: calc.Stack{
				{Type: 0, Value: "V1"},
				{Type: 1, Value: "v2"},
			},
			dest: calc.Stack{},
			expected: calc.Stack{
				{Type: 1, Value: "v2"},
				{Type: 0, Value: "V1"},
			},
		},
		{
			name: "empty one stack into an already filled one",
			source: calc.Stack{
				{Type: 2, Value: "V3"},
				{Type: 3, Value: "v4"},
			},
			dest: calc.Stack{
				{Type: 0, Value: "V1"},
				{Type: 1, Value: "v2"},
			},
			expected: calc.Stack{
				{Type: 0, Value: "V1"},
				{Type: 1, Value: "v2"},
				{Type: 3, Value: "v4"},
				{Type: 2, Value: "V3"},
			},
		},
		{
			name:   "empty nil stack into an already filled one",
			source: nil,
			dest: calc.Stack{
				{Type: 0, Value: "V1"},
				{Type: 1, Value: "v2"},
			},
			expected: calc.Stack{
				{Type: 0, Value: "V1"},
				{Type: 1, Value: "v2"},
			},
		},
		{
			name: "empty stack into an nil one",
			source: calc.Stack{
				{Type: 0, Value: "V1"},
				{Type: 1, Value: "v2"},
			},
			dest: nil,
			expected: calc.Stack{
				{Type: 1, Value: "v2"},
				{Type: 0, Value: "V1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.source.EmptyInto(&tt.dest)

			if len(tt.source) != 0 {
				t.Errorf("source should be empty now, but it is %v", tt.source)
			}

			if !reflect.DeepEqual(tt.dest, tt.expected) {
				t.Errorf("dest should equal the original values of source, but it is %v", tt.dest)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    calc.Stack
		want bool
	}{
		{name: "an empty stack", s: calc.Stack{}, want: true},
		{name: "a nil stack", s: nil, want: true},
		{name: "a filled stack", s: calc.Stack{
			calc.Token{Type: 1, Value: "v1"},
			calc.Token{Type: 2, Value: "v2"},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Length(t *testing.T) {
	tests := []struct {
		name string
		s    calc.Stack
		want int
	}{
		{name: "an empty stack", s: calc.Stack{}, want: 0},
		{name: "a nil stack", s: nil, want: 0},
		{name: "a filled stack", s: calc.Stack{
			calc.Token{Type: 1, Value: "2"},
			calc.Token{Type: 3, Value: "4"},
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name string
		s    calc.Stack
		want calc.Token
	}{
		{name: "an empty stack", s: calc.Stack{}, want: calc.Token{}},
		{name: "a nil stack", s: nil, want: calc.Token{}},
		{name: "a filled stack", s: calc.Stack{
			calc.Token{Type: 1, Value: "2"},
			calc.Token{Type: 3, Value: "4"},
		}, want: calc.Token{Type: 3, Value: "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := copyStack(tt.s)

			if got := tt.s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.s, original) {
				t.Errorf("s should not have changed from %v to %v", original, tt.s)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name       string
		s          calc.Stack
		want       calc.Token
		wantSAfter calc.Stack
	}{
		{name: "an empty stack", s: calc.Stack{}, want: calc.Token{}, wantSAfter: calc.Stack{}},
		{name: "a nil stack", s: nil, want: calc.Token{}, wantSAfter: nil},
		{name: "a filled stack", s: calc.Stack{
			calc.Token{Type: 1, Value: "2"},
			calc.Token{Type: 3, Value: "4"},
		}, want: calc.Token{Type: 3, Value: "4"}, wantSAfter: calc.Stack{
			calc.Token{Type: 1, Value: "2"},
		}},
		{name: "a filled stack with only one element", s: calc.Stack{
			calc.Token{Type: 1, Value: "2"},
		}, want: calc.Token{Type: 1, Value: "2"}, wantSAfter: calc.Stack{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.s, tt.wantSAfter) {
				t.Errorf("the stack should be %v now but it is %v", tt.wantSAfter, tt.s)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name       string
		s          calc.Stack
		tokens     []calc.Token
		wantSAfter calc.Stack
	}{
		{
			name:       "push into empty stack",
			s:          calc.Stack{},
			tokens:     []calc.Token{{Type: 1, Value: "2"}, {Type: 2, Value: "3"}},
			wantSAfter: calc.Stack{{Type: 1, Value: "2"}, {Type: 2, Value: "3"}},
		},
		{
			name:       "push into nil stack",
			s:          nil,
			tokens:     []calc.Token{{Type: 1, Value: "2"}, {Type: 2, Value: "3"}},
			wantSAfter: calc.Stack{{Type: 1, Value: "2"}, {Type: 2, Value: "3"}},
		},
		{
			name:       "push into filled stack",
			s:          calc.Stack{{Type: 5, Value: "5"}, {Type: 6, Value: "6"}},
			tokens:     []calc.Token{{Type: 1, Value: "2"}, {Type: 2, Value: "3"}},
			wantSAfter: calc.Stack{{Type: 5, Value: "5"}, {Type: 6, Value: "6"}, {Type: 1, Value: "2"}, {Type: 2, Value: "3"}},
		},
		{
			name:       "push nil into stack",
			s:          calc.Stack{{Type: 5, Value: "5"}, {Type: 6, Value: "6"}},
			tokens:     nil,
			wantSAfter: calc.Stack{{Type: 5, Value: "5"}, {Type: 6, Value: "6"}},
		},
		{
			name:       "push empty into stack",
			s:          calc.Stack{{Type: 5, Value: "5"}, {Type: 6, Value: "6"}},
			tokens:     calc.Stack{},
			wantSAfter: calc.Stack{{Type: 5, Value: "5"}, {Type: 6, Value: "6"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.tokens...)
			if !reflect.DeepEqual(tt.s, tt.wantSAfter) {
				t.Errorf("s.Push(); s = %v, got %v ", tt.s, tt.wantSAfter)
			}
		})
	}
}
