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
			1,
			2,
		},
		expected: []int{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int{
			1,
			1,
			2,
		},
		expected: []int{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int{
			1,
			2,
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

var uniqueInt8Tests = []struct {
	name     string
	input    []int8
	expected []int8
}{
	{
		name: "no duplicates",
		input: []int8{
			1,
			2,
		},
		expected: []int8{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int8{
			1,
			1,
			2,
		},
		expected: []int8{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int8{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int8{
			1,
			2,
		},
	},
}

func TestUniqueInt8(t *testing.T) {
	for _, tt := range uniqueInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueInt8(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueInt16Tests = []struct {
	name     string
	input    []int16
	expected []int16
}{
	{
		name: "no duplicates",
		input: []int16{
			1,
			2,
		},
		expected: []int16{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int16{
			1,
			1,
			2,
		},
		expected: []int16{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int16{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int16{
			1,
			2,
		},
	},
}

func TestUniqueInt16(t *testing.T) {
	for _, tt := range uniqueInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueInt16(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueInt32Tests = []struct {
	name     string
	input    []int32
	expected []int32
}{
	{
		name: "no duplicates",
		input: []int32{
			1,
			2,
		},
		expected: []int32{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int32{
			1,
			1,
			2,
		},
		expected: []int32{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int32{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int32{
			1,
			2,
		},
	},
}

func TestUniqueInt32(t *testing.T) {
	for _, tt := range uniqueInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueInt32(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueInt64Tests = []struct {
	name     string
	input    []int64
	expected []int64
}{
	{
		name: "no duplicates",
		input: []int64{
			1,
			2,
		},
		expected: []int64{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int64{
			1,
			1,
			2,
		},
		expected: []int64{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int64{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int64{
			1,
			2,
		},
	},
}

func TestUniqueInt64(t *testing.T) {
	for _, tt := range uniqueInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueInt64(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueUintTests = []struct {
	name     string
	input    []uint
	expected []uint
}{
	{
		name: "no duplicates",
		input: []uint{
			1,
			2,
		},
		expected: []uint{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint{
			1,
			1,
			2,
		},
		expected: []uint{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint{
			1,
			2,
		},
	},
}

func TestUniqueUint(t *testing.T) {
	for _, tt := range uniqueUintTests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueUint(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueUint8Tests = []struct {
	name     string
	input    []uint8
	expected []uint8
}{
	{
		name: "no duplicates",
		input: []uint8{
			1,
			2,
		},
		expected: []uint8{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint8{
			1,
			1,
			2,
		},
		expected: []uint8{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint8{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint8{
			1,
			2,
		},
	},
}

func TestUniqueUint8(t *testing.T) {
	for _, tt := range uniqueUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueUint8(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueUint16Tests = []struct {
	name     string
	input    []uint16
	expected []uint16
}{
	{
		name: "no duplicates",
		input: []uint16{
			1,
			2,
		},
		expected: []uint16{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint16{
			1,
			1,
			2,
		},
		expected: []uint16{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint16{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint16{
			1,
			2,
		},
	},
}

func TestUniqueUint16(t *testing.T) {
	for _, tt := range uniqueUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueUint16(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueUint32Tests = []struct {
	name     string
	input    []uint32
	expected []uint32
}{
	{
		name: "no duplicates",
		input: []uint32{
			1,
			2,
		},
		expected: []uint32{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint32{
			1,
			1,
			2,
		},
		expected: []uint32{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint32{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint32{
			1,
			2,
		},
	},
}

func TestUniqueUint32(t *testing.T) {
	for _, tt := range uniqueUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueUint32(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueFloat32Tests = []struct {
	name     string
	input    []float32
	expected []float32
}{
	{
		name: "no duplicates",
		input: []float32{
			1.1,
			2.2,
		},
		expected: []float32{
			1.1,
			2.2,
		},
	},
	{
		name: "one duplicates",
		input: []float32{
			1.1,
			1.1,
			2.2,
		},
		expected: []float32{
			1.1,
			2.2,
		},
	},
	{
		name: "multiple duplicates",
		input: []float32{
			1.1,
			1.1,
			2.2,
			2.2,
			2.2,
			1.1,
			1.1,
			2.2,
			1.1,
		},
		expected: []float32{
			1.1,
			2.2,
		},
	},
}

func TestUniqueFloat32(t *testing.T) {
	for _, tt := range uniqueFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueFloat32(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueFloat64Tests = []struct {
	name     string
	input    []float64
	expected []float64
}{
	{
		name: "no duplicates",
		input: []float64{
			1.1,
			2.2,
		},
		expected: []float64{
			1.1,
			2.2,
		},
	},
	{
		name: "one duplicates",
		input: []float64{
			1.1,
			1.1,
			2.2,
		},
		expected: []float64{
			1.1,
			2.2,
		},
	},
	{
		name: "multiple duplicates",
		input: []float64{
			1.1,
			1.1,
			2.2,
			2.2,
			2.2,
			1.1,
			1.1,
			2.2,
			1.1,
		},
		expected: []float64{
			1.1,
			2.2,
		},
	},
}

func TestUniqueFloat64(t *testing.T) {
	for _, tt := range uniqueFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueFloat64(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueComplex64Tests = []struct {
	name     string
	input    []complex64
	expected []complex64
}{
	{
		name: "no duplicates",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "one duplicates",
		input: []complex64{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "multiple duplicates",
		input: []complex64{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
		},
		expected: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
}

func TestUniqueComplex64(t *testing.T) {
	for _, tt := range uniqueComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueComplex64(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}

var uniqueComplex128Tests = []struct {
	name     string
	input    []complex128
	expected []complex128
}{
	{
		name: "no duplicates",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "one duplicates",
		input: []complex128{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "multiple duplicates",
		input: []complex128{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
		},
		expected: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
}

func TestUniqueComplex128(t *testing.T) {
	for _, tt := range uniqueComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UniqueComplex128(tt.input)
			assert.Equal(t, output, tt.expected)
		})
	}
}
