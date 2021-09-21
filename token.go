package calc

// TokenType defines what the token is.
type TokenType int

// Token represents a pair of a type and a value which matches that type.
type Token struct {
	// Type of the token.
	// This can be used to parse the Value accordingly.
	Type TokenType

	// Value of the Token.
	// It should match the Type.
	Value string
}

// These constants are all possible TokenType values.
const (
	Number TokenType = iota
	Lparen
	Rparen
	Constant
	Function
	Operator
	Whitespace
)
