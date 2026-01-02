package cli

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "		test 	tab     WhiTESPACE		",
			expected: []string{"test", "tab", "whitespace"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("length of cleaned input slice does not match expected\n")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("output word '%s' does not match expected word '%s'", word, expectedWord)
			}
		}
	}
}
