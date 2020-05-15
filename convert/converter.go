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

	// stack := make([]token, 0, len(tokens))
	for _, t := range tokens {
		fmt.Println("Token: ", t)
	}

	return "Not implemented"
}

func tokenize(expression string) []token {
	scanner := bufio.NewScanner(strings.NewReader(expression))

	tokenValues := make([]string, 0, len(expression))
	tokens := make([]token, 0, len(expression))

	split := func(remaining []byte, atEOF bool) (advance int, token []byte, err error) {
		var isNegativeNumber bool
		var startNumberIndex = -1
		for i, b := range remaining {
			if startNumberIndex > -1 && !isNumber(b) {
				if isNegativeNumber {
					return i, remaining[startNumberIndex-1 : i], nil
				}
				return i, remaining[startNumberIndex:i], nil
			} else if isOperator(b) {
				if isNegativeSign(b, tokenValues) {
					isNegativeNumber = true
				} else {
					return i + 1, remaining[i : i+1], nil
				}
			} else if isNumber(b) && startNumberIndex == -1 {
				startNumberIndex = i
			} else if isStartOfGroup(b) || isEndOfGroup(b) {
				return i + 1, remaining[i : i+1], nil
			}
		}

		return len(remaining), nil, nil
	}

	scanner.Split(split)

	for scanner.Scan() {
		t := newToken(scanner.Text())
		if t.tokenType != unknownType {
			tokens = append(tokens, t)
		}
	}

	return tokens
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '/' || b == '*'
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
	if b == '-' && len(currentTokens) > 0 && len(currentTokens[lastItemIndex]) == 1 {
		lastToken := currentTokens[lastItemIndex]
		return !isNumber(lastToken[0])
	}
	return false
}

func (t token) precedence() int {
	if t.tokenType == groupStartType {
		return 3
	}

	if t.tokenType == operatorType {
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
