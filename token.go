package calc

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

const (
	Number TokenType = iota
	Lparen
	Rparen
	Constant
	Function
	Operator
	Whitespace
)
