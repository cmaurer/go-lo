package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var findStringTests = []struct {
	name     string
	input    []string
	equalTo  string
	expected string
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
		equalTo:  "test2",
		expected: "test2",
	},
	{
		name: "more than one element",
		input: []string{
			"test1",
			"test2",
			"test3",
			"test2",
		},
		equalTo:  "test2",
		expected: "test2",
	},
}

func TestFindString(t *testing.T) {
	for _, tt := range findStringTests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindString(tt.input, func(in string) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findIntTests = []struct {
	name     string
	input    []int
	equalTo  int
	expected int
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []int{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindInt(t *testing.T) {
	for _, tt := range findIntTests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindInt(tt.input, func(in int) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findInt8Tests = []struct {
	name     string
	input    []int8
	equalTo  int8
	expected int8
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []int8{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindInt8(t *testing.T) {
	for _, tt := range findInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindInt8(tt.input, func(in int8) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findInt16Tests = []struct {
	name     string
	input    []int16
	equalTo  int16
	expected int16
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []int16{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindInt16(t *testing.T) {
	for _, tt := range findInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindInt16(tt.input, func(in int16) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findInt32Tests = []struct {
	name     string
	input    []int32
	equalTo  int32
	expected int32
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []int32{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindInt32(t *testing.T) {
	for _, tt := range findInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindInt32(tt.input, func(in int32) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findInt64Tests = []struct {
	name     string
	input    []int64
	equalTo  int64
	expected int64
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []int64{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindInt64(t *testing.T) {
	for _, tt := range findInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindInt64(tt.input, func(in int64) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findUintTests = []struct {
	name     string
	input    []uint
	equalTo  uint
	expected uint
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []uint{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindUint(t *testing.T) {
	for _, tt := range findUintTests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindUint(tt.input, func(in uint) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findUint8Tests = []struct {
	name     string
	input    []uint8
	equalTo  uint8
	expected uint8
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []uint8{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindUint8(t *testing.T) {
	for _, tt := range findUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindUint8(tt.input, func(in uint8) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findUint16Tests = []struct {
	name     string
	input    []uint16
	equalTo  uint16
	expected uint16
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []uint16{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindUint16(t *testing.T) {
	for _, tt := range findUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindUint16(tt.input, func(in uint16) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findUint32Tests = []struct {
	name     string
	input    []uint32
	equalTo  uint32
	expected uint32
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
		equalTo:  2,
		expected: 2,
	},
	{
		name: "more than one element",
		input: []uint32{
			1,
			2,
			3,
			2,
		},
		equalTo:  2,
		expected: 2,
	},
}

func TestFindUint32(t *testing.T) {
	for _, tt := range findUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindUint32(tt.input, func(in uint32) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findFloat32Tests = []struct {
	name     string
	input    []float32
	equalTo  float32
	expected float32
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
		equalTo:  2.2,
		expected: 2.2,
	},
	{
		name: "more than one element",
		input: []float32{
			1.1,
			2.2,
			3.3,
			2.2,
		},
		equalTo:  2.2,
		expected: 2.2,
	},
}

func TestFindFloat32(t *testing.T) {
	for _, tt := range findFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindFloat32(tt.input, func(in float32) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findFloat64Tests = []struct {
	name     string
	input    []float64
	equalTo  float64
	expected float64
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
		equalTo:  2.2,
		expected: 2.2,
	},
	{
		name: "more than one element",
		input: []float64{
			1.1,
			2.2,
			3.3,
			2.2,
		},
		equalTo:  2.2,
		expected: 2.2,
	},
}

func TestFindFloat64(t *testing.T) {
	for _, tt := range findFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindFloat64(tt.input, func(in float64) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findComplex64Tests = []struct {
	name     string
	input    []complex64
	equalTo  complex64
	expected complex64
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
		},
		equalTo:  (4.8 + 3.14i),
		expected: (4.8 + 3.14i),
	},
	{
		name: "more than one element",
		input: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
			(4.8 + 3.14i),
		},
		equalTo:  (4.8 + 3.14i),
		expected: (4.8 + 3.14i),
	},
}

func TestFindComplex64(t *testing.T) {
	for _, tt := range findComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindComplex64(tt.input, func(in complex64) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}

var findComplex128Tests = []struct {
	name     string
	input    []complex128
	equalTo  complex128
	expected complex128
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
		},
		equalTo:  (4.8 + 3.14i),
		expected: (4.8 + 3.14i),
	},
	{
		name: "more than one element",
		input: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
			(4.8 + 3.14i),
		},
		equalTo:  (4.8 + 3.14i),
		expected: (4.8 + 3.14i),
	},
}

func TestFindComplex128(t *testing.T) {
	for _, tt := range findComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindComplex128(tt.input, func(in complex128) bool {
				return in == tt.equalTo
			})
			assert.Equal(t, tt.expected, *output)
		})
	}

}
