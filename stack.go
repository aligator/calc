package calc

// Stack is a LIFO data structure.
type Stack []Token

// Pop removes the token at the top of the stack and returns its value.
func (s *Stack) Pop() Token {
	if s.IsEmpty() {
		return Token{}
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element

}

// Push adds tokens to the top of the stack.
func (s *Stack) Push(i ...Token) {
	*s = append(*s, i...)
}

// Peek returns the token at the top of the stack.
func (s *Stack) Peek() Token {
	if s.IsEmpty() {
		return Token{}
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	return element
}

// EmptyInto dumps all tokens from one stack to another.
func (s *Stack) EmptyInto(dest *Stack) {
	if !s.IsEmpty() {
		for i := s.Length() - 1; i >= 0; i-- {
			dest.Push(s.Pop())
		}
	}
}

// IsEmpty checks if there are any tokens in the stack.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Length returns the amount of tokens in the stack.
func (s *Stack) Length() int {
	return len(*s)
}
