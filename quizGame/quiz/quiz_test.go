package quiz

import (
	"reflect"
	"testing"
)

type convertToProblemsCase struct {
	name     string
	lines    [][]string
	expected []Problem
}

func TestConvertToProblems(t *testing.T) {
	normal := convertToProblemsCase{
		name: "Normal input",
		lines: [][]string{
			{"q1", "a1"},
			{"q2", "a2"},
			{"q3", "a3"},
		},
		expected: []Problem{
			{Question: "q1", Answer: "a1"},
			{Question: "q2", Answer: "a2"},
			{Question: "q3", Answer: "a3"},
		},
	}

	hasSpaces := convertToProblemsCase{
		name: "Normal input",
		lines: [][]string{
			{"q1", "   a1"},
			{"q2", "a2   "},
			{"q3", "  a3  "},
		},
		expected: []Problem{
			{Question: "q1", Answer: "a1"},
			{Question: "q2", Answer: "a2"},
			{Question: "q3", Answer: "a3"},
		},
	}

	tests := []convertToProblemsCase{normal, hasSpaces}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ConvertToProblems(test.lines)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected: %#v, Actual: %#v", test.expected, actual)
			}
		})
	}
}
