package main

import(
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},{
			input: " Test case ",
			expected: []string{"test", "case"},
		},
		// add more cases here
	}

	for _, c:= range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("The words not match %s, %s", word, expectedWord)
				return
			}
		}
	}
}

