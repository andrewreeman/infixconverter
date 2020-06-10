package convert

import "testing"

func TestCanConvertCorrectStatements(assert *testing.T) {
	tests := []testPackage{
		// testPackage{
		// 	input:    "44 - -4 +     0  * (4 + 37)",
		// 	expected: "44 -4 - 0 4 37 + * +",
		// },
		// testPackage{
		// 	input:    "-1",
		// 	expected: "-1",
		// },
		// testPackage{
		// 	input:    "456",
		// 	expected: "456",
		// },
		// testPackage{
		// 	input:    "(-1+84)/(5*4)",
		// 	expected: "-1 84 + 5 4 * /",
		// },
		// testPackage{
		// 	input:    "-(-1+84)/(5*4)",
		// 	expected: "-1 84 + 5 4 * / -",
		// },
		// testPackage{
		// 	input:    "4*((2+6)/3)-2^2^3",
		// 	expected: "4 2 6 + 3 / * 2 2 3 ^ ^ -",
		// },
		// testPackage{
		// 	input:    "2^5^3",
		// 	expected: "2 5 3 ^ ^",
		// },
		testPackage{
			input:    "2^5^3+1",
			expected: "2 5 3 ^ ^ 1 +",
		},
		// testPackage{
		// 	input:    "2^5+1",
		// 	expected: "2 5 ^ 1 +",
		// },
		// testPackage{
		// 	input:    "2^5^8^9",
		// 	expected: "2 5 8 9 ^ ^ ^",
		// },
		// testPackage{
		// 	input:    "2^5^(8-4)^9",
		// 	expected: "2 5 8 4 - 9 ^ ^ ^",
		// },
		// testPackage{
		// 	input:    "2^5^8^9+84",
		// 	expected: "2 5 8 9 ^ ^ ^ 84 +",
		// },
	}

	for _, test := range tests {
		result := Convert(test.input)

		if result != test.expected {
			assert.Errorf("Result for input %s did not match %s. Instead we received %s", test.input, test.expected, result)
		}
	}
}

type testPackage struct {
	input    string
	expected string
}
