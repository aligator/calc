package calc

import (
	"fmt"
)

func ShuntingYard(s Stack) (Stack, error) {
	parenthesisCounter := 0

	postfix := Stack{}
	operators := Stack{}
	for _, v := range s {
		switch v.Type {
		case OPERATOR:
			for !operators.IsEmpty() {
				val := v.Value
				top := operators.Peek().Value
				if (oprData[val].prec <= oprData[top].prec && oprData[val].rAsoc == false) ||
					(oprData[val].prec < oprData[top].prec && oprData[val].rAsoc == true) {
					postfix.Push(operators.Pop())
					continue
				}
				break
			}
			operators.Push(v)
		case LPAREN:
			parenthesisCounter++
			operators.Push(v)
		case RPAREN:
			for i := operators.Length() - 1; i >= 0; i-- {
				if operators[i].Type != LPAREN {
					postfix.Push(operators.Pop())
					continue
				} else {
					operators.Pop()
					break
				}
			}
			parenthesisCounter--
		default:
			postfix.Push(v)
		}
	}
	operators.EmptyInto(&postfix)

	if parenthesisCounter != 0 {
		return postfix, fmt.Errorf("invalid parenthesis count - %d do not match", parenthesisCounter)
	}
	return postfix, nil
}
