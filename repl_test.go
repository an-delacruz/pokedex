package main

import (
	"testing"
)


func TestCleanInput(t *testing.T){
	cases := [] struct {
		input string
		expected []string
	}{
		{
			input: "   ",
			expected: []string{},
		},
		{
			input: "   hello   ",
			expected: []string{"hello"},
		},
		{
			input: "   Hello World",
			expected: []string{"hello","world"},
		},
		{
			input: "   hello world",
			expected: []string{"hello","world"},
		},
	}

	for _,c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("lenghts don't match: '%v' vs '%v'",actual,c.expected)
			 continue
		}
		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("cleanInput(%v) == %v, expectec %v", c.input, actual, c.expected)
			}
		}
	}
}