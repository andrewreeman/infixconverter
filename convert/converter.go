package convert

import (
	"bufio"
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

	tokenize(expression)

	return expression
}
func tokenize(expression string) {
	scanner := bufio.NewScanner(strings.NewReader(expression))

	// for {
	// 	if reader.Len() == 0 {
	// 		fmt.Println("Finished reading")
	// 		break
	// 	}

	// 	readRune, length, readError := reader.ReadRune()

	// 	if length == 0 || readError != nil {
	// 		fmt.Println("Unexpectadly finished reading")
	// 		break
	// 	}
	// 	fmt.Println("Read rune ", readRune)
	// }

}
