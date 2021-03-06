package calc

import (
	"errors"
	"io"
)

type tokenBuffer struct {
	tok Token
	n   int
}

// TokenScanner defines a scanner which returns one token on each scan.
type TokenScanner interface {
	// Scan returns one token on each scan.
	// If the previous token was the last one, it must return io.EOF.
	Scan() (Token, error)
}

type Parser struct {
	s   TokenScanner
	buf tokenBuffer
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Scan() (Token, error) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, nil
	}

	tok, err := p.s.Scan()
	if err != nil {
		return Token{}, err
	}

	p.buf.tok = tok

	return tok, nil
}

func (p *Parser) ScanIgnoreWhitespace() (Token, error) {
	tok, err := p.Scan()
	if err != nil {
		return Token{}, err
	}

	if tok.Type == Whitespace {
		return p.Scan()
	}
	return tok, nil
}

func (p *Parser) Unscan() {
	p.buf.n = 1
}

func (p *Parser) Parse() (Stack, error) {
	stack := Stack{}
	for {
		tok, err := p.ScanIgnoreWhitespace()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return Stack{}, err
		} else if tok.Type == Operator && tok.Value == "-" {
			lastTok := stack.Peek()
			nextTok, err := p.ScanIgnoreWhitespace()
			if err != nil {
				return Stack{}, err
			}

			if (lastTok.Type == Operator || lastTok.Value == "" || lastTok.Type == Lparen) && nextTok.Type == Number {
				stack.Push(Token{Number, "-" + nextTok.Value})
			} else {
				stack.Push(tok)
				p.Unscan()
			}
		} else {
			stack.Push(tok)
		}
	}
	return stack, nil
}
