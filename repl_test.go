package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world ",
			expected: []string{"hello","world"},
		},
		{
			input: "   somethinghere is offman",
			expected: []string{"somethinghere","is","offman"},
		},
		{
			input: "I have a normal sentence here",
			expected: []string{"i","have","a","normal","sentence","here"},
		},
		{
			input: "ThIS IS SOMeTHING ",
			expected: []string{"this","is","something"},
		},
		{
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Length missmatch")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("Output missmatch: actual: %s; expected: %s",word,expectedWord)
			}
		}
	}

}