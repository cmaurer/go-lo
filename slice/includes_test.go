package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var includesStringTests = []struct {
	name     string
	input    []string
	included string
	expected bool
}{
	{
		name: "should be true",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
		included: "test2",
		expected: true,
	},
	{
		name: "should be false",
		input: []string{
			"test1",
			"test2",
		},
		included: "test3",
		expected: false,
	},
}

func TestIncludesString(t *testing.T) {
	for _, tt := range includesStringTests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesString(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesIntTests = []struct {
	name     string
	input    []int
	included int
	expected bool
}{
	{
		name: "should be true",
		input: []int{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []int{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesInt(t *testing.T) {
	for _, tt := range includesIntTests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesInt(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesInt8Tests = []struct {
	name     string
	input    []int8
	included int8
	expected bool
}{
	{
		name: "should be true",
		input: []int8{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []int8{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesInt8(t *testing.T) {
	for _, tt := range includesInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesInt8(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesInt16Tests = []struct {
	name     string
	input    []int16
	included int16
	expected bool
}{
	{
		name: "should be true",
		input: []int16{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []int16{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesInt16(t *testing.T) {
	for _, tt := range includesInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesInt16(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesInt32Tests = []struct {
	name     string
	input    []int32
	included int32
	expected bool
}{
	{
		name: "should be true",
		input: []int32{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []int32{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesInt32(t *testing.T) {
	for _, tt := range includesInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesInt32(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesInt64Tests = []struct {
	name     string
	input    []int64
	included int64
	expected bool
}{
	{
		name: "should be true",
		input: []int64{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []int64{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesInt64(t *testing.T) {
	for _, tt := range includesInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesInt64(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesUintTests = []struct {
	name     string
	input    []uint
	included uint
	expected bool
}{
	{
		name: "should be true",
		input: []uint{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []uint{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesUint(t *testing.T) {
	for _, tt := range includesUintTests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesUint(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesUint8Tests = []struct {
	name     string
	input    []uint8
	included uint8
	expected bool
}{
	{
		name: "should be true",
		input: []uint8{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []uint8{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesUint8(t *testing.T) {
	for _, tt := range includesUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesUint8(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesUint16Tests = []struct {
	name     string
	input    []uint16
	included uint16
	expected bool
}{
	{
		name: "should be true",
		input: []uint16{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []uint16{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesUint16(t *testing.T) {
	for _, tt := range includesUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesUint16(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesUint32Tests = []struct {
	name     string
	input    []uint32
	included uint32
	expected bool
}{
	{
		name: "should be true",
		input: []uint32{
			1,
			2,
			3,
		},
		included: 2,
		expected: true,
	},
	{
		name: "should be false",
		input: []uint32{
			1,
			2,
		},
		included: 3,
		expected: false,
	},
}

func TestIncludesUint32(t *testing.T) {
	for _, tt := range includesUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesUint32(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesFloat32Tests = []struct {
	name     string
	input    []float32
	included float32
	expected bool
}{
	{
		name: "should be true",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
		included: 2.2,
		expected: true,
	},
	{
		name: "should be false",
		input: []float32{
			1.1,
			2.2,
		},
		included: 3.3,
		expected: false,
	},
}

func TestIncludesFloat32(t *testing.T) {
	for _, tt := range includesFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesFloat32(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesFloat64Tests = []struct {
	name     string
	input    []float64
	included float64
	expected bool
}{
	{
		name: "should be true",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
		included: 2.2,
		expected: true,
	},
	{
		name: "should be false",
		input: []float64{
			1.1,
			2.2,
		},
		included: 3.3,
		expected: false,
	},
}

func TestIncludesFloat64(t *testing.T) {
	for _, tt := range includesFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesFloat64(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesComplex64Tests = []struct {
	name     string
	input    []complex64
	included complex64
	expected bool
}{
	{
		name: "should be true",
		input: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
		},
		included: (4.8 + 3.14i),
		expected: true,
	},
	{
		name: "should be false",
		input: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
		},
		included: (7.2 + 3.14i),
		expected: false,
	},
}

func TestIncludesComplex64(t *testing.T) {
	for _, tt := range includesComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesComplex64(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}

var includesComplex128Tests = []struct {
	name     string
	input    []complex128
	included complex128
	expected bool
}{
	{
		name: "should be true",
		input: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
		},
		included: (4.8 + 3.14i),
		expected: true,
	},
	{
		name: "should be false",
		input: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
		},
		included: (7.199999999999999 + 3.14i),
		expected: false,
	},
}

func TestIncludesComplex128(t *testing.T) {
	for _, tt := range includesComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
			included := IncludesComplex128(tt.input, tt.included)
			assert.Equal(t, tt.expected, included)
		})
	}

}
