package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var indexOfStringTests = []struct {
	name     string
	input    []string
	element  string
	expected int
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
		element:  "test2",
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []string{
			"test1",
			"test2",
			"test3",
			"test4",
			"test5",
		},
		element:  "test5",
		expected: 4,
	},
	{
		name: "not found",
		input: []string{
			"test1",
			"test2",
			"test3",
			"test4",
			"test5",
		},
		element:  "test10",
		expected: -1,
	},
}

func TestIndexOfString(t *testing.T) {
	for _, tt := range indexOfStringTests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfString(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}
