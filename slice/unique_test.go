package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var uniqueStringTests = []struct {
	name     string
	input    []string
	expected []string
}{
	{
		name: "no duplicates",
		input: []string{
			"test1",
			"test2",
		},
		expected: []string{
			"test1",
			"test2",
		},
	},
	{
		name: "one duplicates",
		input: []string{
			"test1",
			"test1",
			"test2",
		},
		expected: []string{
			"test1",
			"test2",
		},
	},
	{
		name: "multiple duplicates",
		input: []string{
			"test1",
			"test1",
			"test2",
			"test2",
			"test2",
			"test1",
			"test1",
			"test2",
			"test1",
		},
		expected: []string{
			"test1",
			"test2",
		},
	},
}

func TestUniqueString(t *testing.T) {

	for _, tt := range uniqueStringTests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueString(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}

var uniqueIntTests = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name: "no duplicates",
		input: []int{
			0,
			1,
		},
		expected: []int{
			0,
			1,
		},
	},
	{
		name: "one duplicates",
		input: []int{
			1,
			1,
			3,
		},
		expected: []int{
			1,
			3,
		},
	},
	{
		name: "multiple duplicates",
		input: []int{
			1,
			2,
			3,
			3,
			3,
			1,
			1,
			1,
			3,
		},
		expected: []int{
			1,
			2,
			3,
		},
	},
}

func TestUniqueInt(t *testing.T) {

	for _, tt := range uniqueIntTests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueInt(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}
