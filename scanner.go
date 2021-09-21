package calc

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Read() (rune, error) {
	ch, _, err := s.r.ReadRune()
	return ch, err
}

func (s *Scanner) Unread() error {
	return s.r.UnreadRune()
}

func (s *Scanner) loadNextRuneTo(buf *bytes.Buffer) error {
	r, err := s.Read()
	if err != nil {
		return err
	}
	buf.WriteRune(r)

	return nil
}

func (s *Scanner) Scan() (Token, error) {
	ch, err := s.Read()
	if err != nil {
		return Token{}, err
	}

	if unicode.IsDigit(ch) {
		err = s.Unread()
		if err != nil {
			return Token{}, err
		}

		return s.ScanNumber()
	} else if unicode.IsLetter(ch) {
		err = s.Unread()
		if err != nil {
			return Token{}, err
		}

		return s.ScanWord()
	} else if IsOperator(ch) {
		return Token{Operator, string(ch)}, nil
	} else if unicode.IsSpace(ch) {
		err = s.Unread()
		if err != nil {
			return Token{}, err
		}

		return s.ScanWhitespace()
	}

	switch ch {
	case '(':
		return Token{Lparen, "("}, nil
	case ')':
		return Token{Rparen, ")"}, nil
	}

	return Token{}, fmt.Errorf("invalid token %v", ch)
}

func (s *Scanner) ScanWord() (Token, error) {
	var buf bytes.Buffer
	if err := s.loadNextRuneTo(&buf); err != nil {
		return Token{}, err
	}

	for {
		if ch, err := s.Read(); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return Token{}, err
		} else if ch == '(' {
			_, err = buf.WriteRune(ch)
			if err != nil {
				return Token{}, err
			}

			parentCount := 1
			for parentCount > 0 {
				fch, err := s.Read()
				if err != nil {
					return Token{}, err
				}

				if fch == '(' {
					parentCount += 1
					_, err = buf.WriteRune(fch)
					if err != nil {
						return Token{}, err
					}
				} else if fch == ')' {
					parentCount -= 1
					_, err = buf.WriteRune(fch)
					if err != nil {
						return Token{}, err
					}
				} else {
					_, err = buf.WriteRune(fch)
					if err != nil {
						return Token{}, err
					}
				}
			}
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			err = s.Unread()
			if err != nil {
				return Token{}, err
			}
			break
		} else {
			_, err = buf.WriteRune(ch)
			if err != nil {
				return Token{}, err
			}
		}
	}

	value := strings.ToUpper(buf.String())
	if strings.ContainsAny(value, "()") {
		return Token{Function, value}, nil
	} else {
		return Token{Constant, value}, nil
	}
}

func (s *Scanner) ScanNumber() (Token, error) {
	var buf bytes.Buffer
	if err := s.loadNextRuneTo(&buf); err != nil {
		return Token{}, err
	}

	for {
		if ch, err := s.Read(); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return Token{}, err
		} else if !unicode.IsDigit(ch) && ch != '.' {
			err = s.Unread()
			if err != nil {
				return Token{}, err
			}
			break
		} else {
			_, err = buf.WriteRune(ch)
			if err != nil {
				return Token{}, err
			}
		}
	}

	return Token{Number, buf.String()}, nil
}

func (s *Scanner) ScanWhitespace() (Token, error) {
	var buf bytes.Buffer
	if err := s.loadNextRuneTo(&buf); err != nil {
		return Token{}, err
	}

	for {
		if ch, err := s.Read(); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return Token{}, err
		} else if !unicode.IsSpace(ch) {
			err = s.Unread()
			if err != nil {
				return Token{}, err
			}
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Whitespace, buf.String()}, nil
}

func IsOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '^'
}
