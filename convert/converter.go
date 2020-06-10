package convert

import (
	"bufio"
	"fmt"
	"strings"
)

// Convert the provided expression from infix to a postfix notation
func Convert(expression string) string {
	if len(expression) == 0 {
		return expression
	}

	tokens := tokenize(expression)
	tokens = toPostFix(tokens, nil, 0)

	builder := strings.Builder{}
	builder.Grow(len(tokens))

	for _, t := range tokens {
		builder.WriteString(t.value + " ")
	}

	return strings.TrimSpace(builder.String())
}

func toPostFix(tokens []token, rightAssociativityToken *token, depth int) []token {
	fmt.Println("===Post fix called at depth===", depth)
	stack := make([]token, 0, len(tokens))
	tmpToken := newUnknown()
	tokensCount := len(tokens)
	for i, t := range tokens {
		fmt.Println("Processing token ", t.value)
		fmt.Println("depth ", depth)
		if t.tokenType == numericType {
			fmt.Println("Current token is number")
			fmt.Println("depth ", depth)
			if tmpToken.tokenType == operatorType {
				fmt.Println("Buffered token is operator")
				if i < (tokensCount - 1) {
					nextToken := tokens[i+1]
					if nextToken.precedence() > tmpToken.precedence() {

						fmt.Println("Token is higher precedence than buffered token", nextToken.value, tmpToken.value)
						fmt.Println("depth ", depth)

						rightOperandExpressionStack := toPostFix(tokens[i:], rightAssociativityToken, depth+1)
						stack = append(stack, rightOperandExpressionStack...)
						stack = append(stack, tmpToken)
						return stack
					} else if rightAssociativityToken != nil && (*rightAssociativityToken).value == "^" && nextToken.tokenType == operatorType && nextToken.precedence() < (*rightAssociativityToken).precedence() {
						fmt.Println("BOOM!", nextToken)
						fmt.Println("Stack when found token of lower precedence than right asoc")
						fmt.Println("depth ", depth)

						for _, _t := range stack {
							fmt.Println(_t.value)
						}

						fmt.Println("Now current will be left operand. Now obtaining right operand expression at depth", depth)
						rightOperandExpressionStack := toPostFix(tokens[i+1:], nil, depth+1)
						fmt.Println("Printing right operand expression at depth ", depth)
						for _, _t := range rightOperandExpressionStack {
							fmt.Println(_t.value)
						}
						stack = append(stack, t)
						stack = append(stack, tmpToken)
						stack = append(stack, *rightAssociativityToken)
						stack = append(stack, rightOperandExpressionStack...)

						fmt.Println("Printing stack after adding tokens at depth ", depth)
						for _, _t := range stack {
							fmt.Println(_t.value)
						}
						return stack

					} else if nextToken.value[0] == '^' && tmpToken.value[0] == '^' {
						fmt.Println("Current token and buffered token are both ^")
						fmt.Println("depth ", depth)
						fmt.Println("Current stack is ")
						for _, _t := range stack {
							fmt.Println(_t.value)
						}
						rightOperandExpressionStack := toPostFix(tokens[i:], &tmpToken, depth+1)
						stack = append(stack, rightOperandExpressionStack...)

						fmt.Println("index is ", i)

						fmt.Println("Stack is after ")
						fmt.Println("depth ", depth)
						for _, _t := range stack {
							fmt.Println(_t.value)
						}
						fmt.Println("Finished printing stack")
						fmt.Println("Current token is ", t)
						fmt.Println("depth ", depth)
						return stack
					}
				}
				fmt.Println("Adding token and tmp token to stack", t.value, tmpToken.value)
				fmt.Println("depth ", depth)
				stack = append(stack, t)
				stack = append(stack, tmpToken)

				if rightAssociativityToken != nil && (*rightAssociativityToken).value == "^" {
					stack = append(stack, *rightAssociativityToken)
				}

				tmpToken.tokenType = unknownType
			} else {
				fmt.Println("Adding token to stack")
				fmt.Println("depth ", depth)
				stack = append(stack, t)
			}

		} else if t.tokenType == operatorType {
			fmt.Println("Current token is operator")
			fmt.Println("depth ", depth)
			tmpToken = t
		} else if t.tokenType == groupStartType && tmpToken.tokenType == operatorType {
			rightOperandExpressionStack := toPostFix(tokens[i+1:], rightAssociativityToken, depth+1)
			stack = append(stack, rightOperandExpressionStack...)
			stack = append(stack, tmpToken)
			return stack
		}

	}
	return stack
}

func tokenize(expression string) []token {
	scanner := bufio.NewScanner(strings.NewReader(expression))

	tokenValues := make([]string, 0, len(expression))
	tokens := make([]token, 0, len(expression))

	split := func(remaining []byte, atEOF bool) (advance int, token []byte, err error) {
		var isNegativeNumber bool
		var startNumberIndex = -1
		for i, b := range remaining {
			if isStartOfGroup(b) && isNegativeNumber {
				return i, remaining[i-1 : i], nil
			} else if startNumberIndex > -1 && !isNumber(b) {
				startNumberIndexTmp := startNumberIndex
				startNumberIndex = -1
				if isNegativeNumber {
					return i, remaining[startNumberIndexTmp-1 : i], nil
				}
				return i, remaining[startNumberIndexTmp:i], nil
			} else if isOperator(b) {
				if isNegativeSign(b, tokenValues) {
					isNegativeNumber = true
				} else {
					startNumberIndex = -1
					return i + 1, remaining[i : i+1], nil
				}
			} else if isNumber(b) && startNumberIndex == -1 {
				startNumberIndex = i
			} else if isStartOfGroup(b) || isEndOfGroup(b) {
				startNumberIndex = -1
				return i + 1, remaining[i : i+1], nil
			}
		}

		if startNumberIndex > -1 {
			if isNegativeNumber {
				return len(remaining), remaining[startNumberIndex-1:], nil
			}
			return len(remaining), remaining[startNumberIndex:], nil
		}
		return len(remaining), nil, nil
	}

	scanner.Split(split)

	for scanner.Scan() {

		tokenValue := strings.TrimSpace(scanner.Text())
		tokenValues = append(tokenValues, tokenValue)
		t := newToken(tokenValue)
		if t.tokenType != unknownType {
			tokens = append(tokens, t)
		}
	}

	return tokens
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '/' || b == '*' || b == '^'
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isStartOfGroup(b byte) bool {
	return b == '('
}

func isEndOfGroup(b byte) bool {
	return b == ')'
}

func isNegativeSign(b byte, currentTokens []string) bool {
	lastItemIndex := len(currentTokens) - 1

	if b == '-' {
		if len(currentTokens) > 0 && len(currentTokens[lastItemIndex]) == 1 {
			lastToken := currentTokens[lastItemIndex]
			return !isNumber(lastToken[0])
		} else if len(currentTokens) == 0 {
			return true
		}
	}
	return false
}

func (t token) precedence() int {
	if t.tokenType == groupStartType {
		return 4
	}

	if t.tokenType == operatorType {
		if t.value == "^" {
			return 3
		}

		if t.value == "*" || t.value == "/" {
			return 2
		}
		return 1
	}

	return 0
}

func (t token) String() string {
	return fmt.Sprintf("Value: %s, Type: %d, Precedence: %d", t.value, t.tokenType, t.precedence())
}

const (
	operatorType = iota
	numericType
	groupStartType
	groupEndType
	unknownType = -1
)

type token struct {
	value     string
	tokenType int
}

func newToken(value string) token {
	l := len(value)
	if l == 1 {
		b := value[0]
		if isOperator(b) {
			return newOperator(value)
		}

		if isStartOfGroup(b) {
			return newGroupStart(value)
		}

		if isEndOfGroup(b) {
			return newGroupEnd(value)
		}
	}

	if isNumber(value[l-1]) {
		return newNumber(value)
	}

	return newUnknown()
}

func newOperator(value string) token {
	return token{value, operatorType}
}

func newNumber(value string) token {
	return token{value, numericType}
}

func newGroupStart(value string) token {
	return token{value, groupStartType}
}

func newGroupEnd(value string) token {
	return token{value, groupEndType}
}

func newUnknown() token {
	return token{"", unknownType}
}
