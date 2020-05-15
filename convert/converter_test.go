package convert

import "testing"

type testPackage struct {
	input    string
	expected string
}

func TestCanConvertCorrectStatements(assert *testing.T) {
	tests := []testPackage{
		testPackage{
			input:    "44 - -4 +     0  * (4 + 37)",
			expected: "44 -4 - 0 4 37 + * +",
		},
	}

	// input := "44 - -4 +     0  * (4 + 37)"
	// result := Convert(input)
	// expected := "44 -4 - 0 4 37 + * +"

	// if result != expected {
	// 	t.Errorf("Result for input %s did not match %s. Instead we received %s", input, expected, result)
	// }

	for _, test := range tests {
		result := Convert(test.input)

		if result != test.expected {
			assert.Errorf("Result for input %s did not match %s. Instead we received %s", test.input, test.expected, result)
		}
	}

}
