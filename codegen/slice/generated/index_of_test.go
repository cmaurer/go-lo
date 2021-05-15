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


var indexOfIntTests = []struct {
	name     string
	input    []int
	element  int
	expected int
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []int{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []int{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfInt(t *testing.T) {
	for _, tt := range indexOfIntTests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfInt(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfInt8Tests = []struct {
	name     string
	input    []int8
	element  int8
	expected int
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []int8{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []int8{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfInt8(t *testing.T) {
	for _, tt := range indexOfInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfInt8(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfInt16Tests = []struct {
	name     string
	input    []int16
	element  int16
	expected int
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []int16{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []int16{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfInt16(t *testing.T) {
	for _, tt := range indexOfInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfInt16(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfInt32Tests = []struct {
	name     string
	input    []int32
	element  int32
	expected int
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []int32{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []int32{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfInt32(t *testing.T) {
	for _, tt := range indexOfInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfInt32(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfInt64Tests = []struct {
	name     string
	input    []int64
	element  int64
	expected int
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []int64{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []int64{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfInt64(t *testing.T) {
	for _, tt := range indexOfInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfInt64(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfUintTests = []struct {
	name     string
	input    []uint
	element  uint
	expected int
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []uint{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []uint{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfUint(t *testing.T) {
	for _, tt := range indexOfUintTests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfUint(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfUint8Tests = []struct {
	name     string
	input    []uint8
	element  uint8
	expected int
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []uint8{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []uint8{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfUint8(t *testing.T) {
	for _, tt := range indexOfUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfUint8(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfUint16Tests = []struct {
	name     string
	input    []uint16
	element  uint16
	expected int
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []uint16{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []uint16{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfUint16(t *testing.T) {
	for _, tt := range indexOfUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfUint16(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfUint32Tests = []struct {
	name     string
	input    []uint32
	element  uint32
	expected int
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
		element:  2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []uint32{
			1,
			2,
			3,
			4,
			5,
		},
		element:  5,
		expected: 4,
	},
	{
		name: "not found",
		input: []uint32{
			1,
			2,
			3,
			4,
			5,
		},
		element:  10,
		expected: -1,
	},
}

func TestIndexOfUint32(t *testing.T) {
	for _, tt := range indexOfUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfUint32(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfFloat32Tests = []struct {
	name     string
	input    []float32
	element  float32
	expected int
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
		element:  2.2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []float32{
			1.1,
			2.2,
			3.3,
			4.4,
			5.5,
		},
		element:  5.5,
		expected: 4,
	},
	{
		name: "not found",
		input: []float32{
			1.1,
			2.2,
			3.3,
			4.4,
			5.5,
		},
		element:  10.10,
		expected: -1,
	},
}

func TestIndexOfFloat32(t *testing.T) {
	for _, tt := range indexOfFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfFloat32(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfFloat64Tests = []struct {
	name     string
	input    []float64
	element  float64
	expected int
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
		element:  2.2,
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []float64{
			1.1,
			2.2,
			3.3,
			4.4,
			5.5,
		},
		element:  5.5,
		expected: 4,
	},
	{
		name: "not found",
		input: []float64{
			1.1,
			2.2,
			3.3,
			4.4,
			5.5,
		},
		element:  10.10,
		expected: -1,
	},
}

func TestIndexOfFloat64(t *testing.T) {
	for _, tt := range indexOfFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfFloat64(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfComplex64Tests = []struct {
	name     string
	input    []complex64
	element  complex64
	expected int
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
		element:  (4.8+3.14i),
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
			(9.6+3.14i),
			(12+3.14i),
		},
		element:  (12+3.14i),
		expected: 4,
	},
	{
		name: "not found",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
			(9.6+3.14i),
			(12+3.14i),
		},
		element:  (24+3.14i),
		expected: -1,
	},
}

func TestIndexOfComplex64(t *testing.T) {
	for _, tt := range indexOfComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfComplex64(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}


var indexOfComplex128Tests = []struct {
	name     string
	input    []complex128
	element  complex128
	expected int
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
		element:  (4.8+3.14i),
		expected: 1,
	},
	{
		name: "basic example with more elements",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
			(9.6+3.14i),
			(12+3.14i),
		},
		element:  (12+3.14i),
		expected: 4,
	},
	{
		name: "not found",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
			(9.6+3.14i),
			(12+3.14i),
		},
		element:  (24+3.14i),
		expected: -1,
	},
}

func TestIndexOfComplex128(t *testing.T) {
	for _, tt := range indexOfComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := IndexOfComplex128(tt.element, tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}

}
