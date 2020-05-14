package convert

import (
	"bufio"
	"fmt"
	"strings"
)

// type token interface {
// }

// type tokenizer interface {
// }

// type numberToken struct {
// 	token      string
// 	isNegative bool
// }

// type numberTokenizer struct {
// 	buffer string
// }

// func (tokenizer numberTokenizer) tokenize(reader *strings.Reader) *numberToken {
// 	strconv.ParseFloat(character)
// 	return nil
// }

// Convert the provided expression from infix to a postfix notation
func Convert(expression string) string {
	if len(expression) == 0 {
		return expression
	}

	tokenize(expression)

	return "Not implemented"
}
func tokenize(expression string) {
	scanner := bufio.NewScanner(strings.NewReader(expression))

	tokens := make([]string, 0, len(expression))

	split := func(remaining []byte, atEOF bool) (advance int, token []byte, err error) {
		var negativeSign byte
		for i, b := range remaining {
			fmt.Println("Byte: ", b)

			if isOperator(b) {
				if isNegativeSign(b, tokens) {
					negativeSign = b
				} else {
					return i + 1, remaining[i : i+1], nil
				}
			} else if isNumber(b) {
				if negativeSign == '-' {
					return i + 1, remaining[i-1 : i+1], nil
				}
				return i + 1, remaining[i : i+1], nil
			} else if b == '(' || b == ')' {
				return i + 1, remaining[i : i+1], nil
			}
		}

		return len(remaining), nil, nil
	}

	scanner.Split(split)

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	for _, t := range tokens {
		fmt.Println("Token: ", t)
	}
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '/' || b == '*'
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isNegativeSign(b byte, currentTokens []string) bool {
	lastItemIndex := len(currentTokens) - 1
	if b == '-' && len(currentTokens) > 0 && len(currentTokens[lastItemIndex]) == 1 {
		lastToken := currentTokens[lastItemIndex]
		return !isNumber(lastToken[0])
	}
	return false
}
