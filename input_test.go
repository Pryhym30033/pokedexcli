package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"extra space": {
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		"seperated without space": {
			input:    "hello..world ",
			expected: []string{"hello..world"},
		},
		"space between each letter": {
			input:    "h e l l o W o r l d",
			expected: []string{"h", "e", "l", "l", "o", "W", "o", "r", "l", "d"},
		},
		"random space": {
			input:    "ab c",
			expected: []string{"ab", "c"},
		},
	}

	for name, c := range cases {
		got := cleantInput(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("%s has a output size mis-match: got %v : expected %v", name, len(got), len(c.expected))
		}
		for i := range got {
			word := got[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("the output \"%v\" in \"%s\" does not match the expected output \"%v\"", word, name, expectedWord)
			}
		}
	}

}
