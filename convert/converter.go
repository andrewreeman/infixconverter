package convert

import (
	"bufio"
	"fmt"
	"strings"
)

type token string

// Convert the provided expression from infix to a postfix notation
func Convert(expression string) string {
	if len(expression) == 0 {
		return expression
	}

	tokens := tokenize(expression)
	for _, t := range tokens {
		fmt.Println("Token: ", t)
	}

	return "Not implemented"
}

func tokenize(expression string) []token {
	scanner := bufio.NewScanner(strings.NewReader(expression))

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
				if isNegativeSign(b, tokens) {
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
		tokens = append(tokens, token(scanner.Text()))
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

func isNegativeSign(b byte, currentTokens []token) bool {
	lastItemIndex := len(currentTokens) - 1
	if b == '-' && len(currentTokens) > 0 && len(currentTokens[lastItemIndex]) == 1 {
		lastToken := currentTokens[lastItemIndex]
		return !isNumber(lastToken[0])
	}
	return false
}
