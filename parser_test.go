package calc

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"
)

type tokenOrErr struct {
	token Token
	err   error
}

type tokenOrErrStack []tokenOrErr

func (t tokenOrErrStack) toStack() Stack {
	res := Stack{}
	for _, token := range t {
		if token.err != io.EOF {
			res.Push(token.token)
		}
	}

	return res
}

type fakeScanner struct {
	results tokenOrErrStack
	current int
}

func (f *fakeScanner) Scan() (Token, error) {
	tokenOrErr := f.results[f.current]
	f.current++

	if tokenOrErr.err != nil {
		return Token{}, tokenOrErr.err
	}
	return tokenOrErr.token, nil
}

func newFakeScanner(tokens tokenOrErrStack) *fakeScanner {
	return &fakeScanner{results: tokens}
}

var (
	testTokensNormal = tokenOrErrStack{
		{token: Token{Type: Number, Value: "42"}},
		{token: Token{Type: Constant, Value: "PI"}},
		{token: Token{Type: Function, Value: "COS"}},
		{err: io.EOF},
	}

	testTokensEmpty = tokenOrErrStack{{err: io.EOF}}

	testTokensWithError = tokenOrErrStack{
		{token: Token{Type: Number, Value: "42"}},
		{token: Token{Type: Constant, Value: "PI"}},
		{err: errors.New("some error")},
		{token: Token{Type: Function, Value: "COS"}},
		{err: io.EOF},
	}
)

func TestParser_Parse(t *testing.T) {
	type fields struct {
		s   TokenScanner
		buf tokenBuffer
	}
	tests := []struct {
		name    string
		fields  fields
		want    Stack
		wantErr bool
	}{
		{
			name: "a simple stack",
			fields: fields{
				s: newFakeScanner(testTokensNormal),
			},
			want:    testTokensNormal.toStack(),
			wantErr: false,
		},
		{
			name: "an empty stack",
			fields: fields{
				s: newFakeScanner(testTokensEmpty),
			},
			want:    testTokensEmpty.toStack(),
			wantErr: false,
		},
		{
			name: "some error while scanning",
			fields: fields{
				s: newFakeScanner(testTokensWithError),
			},
			want:    Stack{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				s:   tt.fields.s,
				buf: tt.fields.buf,
			}
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Scan(t *testing.T) {
	type fields struct {
		s   *Scanner
		buf struct {
			tok Token
			n   int
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    Token
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				s:   tt.fields.s,
				buf: tt.fields.buf,
			}
			got, err := p.Scan()
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ScanIgnoreWhitespace(t *testing.T) {
	type fields struct {
		s   *Scanner
		buf struct {
			tok Token
			n   int
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				s:   tt.fields.s,
				buf: tt.fields.buf,
			}
			got, err := p.ScanIgnoreWhitespace()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScanIgnoreWhitespace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScanIgnoreWhitespace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Unscan(t *testing.T) {
	type fields struct {
		s   *Scanner
		buf struct {
			tok Token
			n   int
		}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				s:   tt.fields.s,
				buf: tt.fields.buf,
			}

			fmt.Println(p)
		})
	}
}
