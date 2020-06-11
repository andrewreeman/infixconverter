package convert

import (
	"bufio"
	"fmt"
	"strings"
)

// //Convert2 follows the tutorial here on converting infix to postfix: https://www.tutorialspoint.com/Convert-Infix-to-Postfix-Expression
// func Convert2(expression string) string {
// 	stack := make([]string, 0, len(expression))

// 	for _, c := range expression {
// 		stack = append(stack, string(c))
// 	}

// 	for {
// 		pChar, popped := pop(stack)
// 		if pChar == nil {
// 			fmt.Println("Stack is now empty")
// 			break
// 		}

// 		stack = popped
// 		fmt.Println(*pChar)
// 	}

// 	return expression
// }

// Convert2 will convert the provided expression from infix to a postfix notation
func Convert2(expression string) string {
	if len(expression) == 0 {
		return expression
	}

	tokens := tokenize(expression)
	tokens = toPostFix(tokens)

	builder := strings.Builder{}
	builder.Grow(len(tokens))

	for _, t := range tokens {
		builder.WriteString(t.value + " ")
	}

	return strings.TrimSpace(builder.String())
}

func toPostFix(tokens []token) []token {
	fmt.Println("Converting to post fix", tokens)

	postFix := make([]token, 0, len(tokens))
	workingStack := make([]token, 0, len(tokens))

	stackUnderflowProtectionChar := "#"
	workingStack = append(workingStack, newToken(stackUnderflowProtectionChar))

	for _, t := range tokens {
		fmt.Println("Processing token", t)
		if t.tokenType == numericType {
			postFix = append(postFix, t)
		} else if t.tokenType == groupStartType {
			workingStack = append(workingStack, t)
		} else if t.value == "^" {
			workingStack = append(workingStack, t)
		} else if t.tokenType == groupEndType {
			for {
				top := workingStack[len(workingStack)-1]
				fmt.Println("Top value on stack when processing group is", top)
				if top.value == stackUnderflowProtectionChar || top.tokenType == groupStartType {
					fmt.Println("Processing group and reached token", top)
					break
				}

				pChar, popped := pop(workingStack)
				if pChar == nil {
					fmt.Println("Stack is now empty. Unexpected")
					break
				}
				workingStack = popped

				postFix = append(postFix, *pChar)

				fmt.Println(*pChar)
			}

			fmt.Println("pop ( from the stack")
			poppedChar, poppedStack := pop(workingStack)
			fmt.Println("Popped char is ", *poppedChar)
			workingStack = poppedStack
		} else {

			if len(workingStack) == 0 {
				fmt.Println("Working stack is zero for token", t)
				workingStack = append(workingStack, t)
			} else {
				top := workingStack[len(workingStack)-1]
				if t.precedence() > top.precedence() {
					fmt.Println("Token has higher precedence than top so adding to stack", t, top)
					workingStack = append(workingStack, t)
				} else {
					fmt.Println("Popping stack until higher precedence is found")
					for {
						top = workingStack[len(workingStack)-1]
						if top.value == stackUnderflowProtectionChar || t.precedence() > top.precedence() {
							fmt.Println("Popping stack stopped due to ", t, top)
							break
						}

						postFix = append(postFix, top)
						pToken, poppedStack := pop(workingStack)

						workingStack = poppedStack
						if pToken == nil {
							fmt.Println("Stack is now empty. Unexpected")
							break
						}
					}

					workingStack = append(workingStack, t)
				}
			}
		}
	} // for t in tokens

	if len(workingStack) > 0 {
		for {
			top := workingStack[len(workingStack)-1]
			if top.value == stackUnderflowProtectionChar {
				break
			}

			postFix = append(postFix, top)
			pToken, poppedStack := pop(workingStack)

			workingStack = poppedStack
			if pToken == nil {
				fmt.Println("Stack is now empty. Unexpected")
				break
			}
		}
	}

	return postFix
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
			fmt.Println("Checking for negative number. Preceding is", lastToken)
			return lastToken != ")" && !isNumber(lastToken[0])
		} else if len(currentTokens) == 0 {
			return true
		}
	}
	return false
}

func (t token) precedence() int {
	// if t.tokenType == groupStartType {
	// 	return 4
	// }

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

	return newUnknown(value)
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

func newUnknown(value string) token {
	return token{value, unknownType}
}

func pop(stack []token) (*token, []token) {
	l := len(stack)
	if l == 0 {
		return nil, stack
	}

	i := l - 1

	return &stack[i], stack[:i]
}
