package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var oprData = map[string]struct {
	prec  int
	rAsoc bool // true = right // false = left
	fx    func(x, y float64) float64
}{
	"^": {4, true, func(x, y float64) float64 { return math.Pow(x, y) }},
	"*": {3, false, func(x, y float64) float64 { return x * y }},
	"/": {3, false, func(x, y float64) float64 { return x / y }},
	"+": {2, false, func(x, y float64) float64 { return x + y }},
	"-": {2, false, func(x, y float64) float64 { return x - y }},
}

var funcs = map[string]func(x float64) float64{
	"LN":    math.Log,
	"ABS":   math.Abs,
	"COS":   math.Cos,
	"SIN":   math.Sin,
	"TAN":   math.Tan,
	"ACOS":  math.Acos,
	"ASIN":  math.Asin,
	"ATAN":  math.Atan,
	"SQRT":  math.Sqrt,
	"CBRT":  math.Cbrt,
	"CEIL":  math.Ceil,
	"FLOOR": math.Floor,
}

var consts = map[string]float64{
	"E":       math.E,
	"PI":      math.Pi,
	"PHI":     math.Phi,
	"SQRT2":   math.Sqrt2,
	"SQRTE":   math.SqrtE,
	"SQRTPI":  math.SqrtPi,
	"SQRTPHI": math.SqrtPhi,
	"LN2":     math.Ln2,
	"LN10":    math.Ln10,
	"LOG2E":   math.Log2E,
	"LOG10E":  math.Log10E,
}

// SolvePostfix evaluates and returns the answer of the expression converted to postfix
func SolvePostfix(tokens Stack) (float64, error) {
	stack := Stack{}
	for _, v := range tokens {
		switch v.Type {
		case Number:
			stack.Push(v)
		case Function:
			res, err := SolveFunction(v.Value)
			if err != nil {
				return 0, err
			}
			stack.Push(Token{Number, res})
		case Constant:
			if val, ok := consts[v.Value]; ok {
				stack.Push(Token{Number, strconv.FormatFloat(val, 'f', -1, 64)})
			}
		case Operator:
			f := oprData[v.Value].fx
			y, err := strconv.ParseFloat(stack.Pop().Value, 64)
			if err != nil {
				return 0, err
			}
			x, err := strconv.ParseFloat(stack.Pop().Value, 64)
			if err != nil {
				return 0, err
			}

			result := f(x, y)
			stack.Push(Token{Number, strconv.FormatFloat(result, 'f', -1, 64)})
		}
	}

	if len(stack) == 0 {
		return 0, errors.New("empty stack - calculation could not be solved")
	}

	return strconv.ParseFloat(stack[0].Value, 64)
}

// SolveFunction returns the answer of a function found within an expression
func SolveFunction(s string) (string, error) {
	var fArg float64
	var err error

	fType := s[:strings.Index(s, "(")]
	args := s[strings.Index(s, "(")+1 : strings.LastIndex(s, ")")]
	if !strings.ContainsAny(args, "+*-/^") && !ContainsLetter(args) {
		fArg, err = strconv.ParseFloat(args, 64)
		if err != nil {
			return "", err
		}
	} else {
		fArg, err = Solve(args)
		if err != nil {
			return "", err
		}
	}

	function, ok := funcs[fType]
	if !ok {
		return "", fmt.Errorf("function does not exist: %s", fType)
	}

	return strconv.FormatFloat(function(fArg), 'f', -1, 64), nil
}

// ContainsLetter checks if a string contains a letter
func ContainsLetter(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

// Solve a mathematical calculation.
func Solve(s string) (float64, error) {
	stack, err := NewParser(strings.NewReader(s)).Parse()
	if err != nil {
		return 0, err
	}

	stack, err = ShuntingYard(stack)
	if err != nil {
		return 0, err
	}

	return SolvePostfix(stack)
}
