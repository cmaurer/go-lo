package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
		
var reduceStringTests = []struct {
	name     string
	input    []string
	expected func() string
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
		expected: func() string {
			input := []string{
				"test0",
				"test1",
				"test2",
				"test3",
			}
			var accumulator string
			for _, in := range input {
			  accumulator = string(accumulator) + string(in)
			}
			return accumulator
		},
	},
}

func TestReduceString(t *testing.T) {
	for _, tt := range reduceStringTests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceString(tt.input, func(acc string, elem string) string {
                return acc + elem
            }, "test0")
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceIntTests = []struct {
	name     string
	input    []int
	expected func() int
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
		expected: func() int {
			input := []int{
				0,
				1,
				2,
				3,
			}
			var accumulator int
			for _, in := range input {
			  accumulator = int(accumulator) + int(in)
			}
			return accumulator
		},
	},
}

func TestReduceInt(t *testing.T) {
	for _, tt := range reduceIntTests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceInt(tt.input, func(acc int, elem int) int {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceInt8Tests = []struct {
	name     string
	input    []int8
	expected func() int8
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
		expected: func() int8 {
			input := []int8{
				0,
				1,
				2,
				3,
			}
			var accumulator int8
			for _, in := range input {
			  accumulator = int8(accumulator) + int8(in)
			}
			return accumulator
		},
	},
}

func TestReduceInt8(t *testing.T) {
	for _, tt := range reduceInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceInt8(tt.input, func(acc int8, elem int8) int8 {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceInt16Tests = []struct {
	name     string
	input    []int16
	expected func() int16
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
		expected: func() int16 {
			input := []int16{
				0,
				1,
				2,
				3,
			}
			var accumulator int16
			for _, in := range input {
			  accumulator = int16(accumulator) + int16(in)
			}
			return accumulator
		},
	},
}

func TestReduceInt16(t *testing.T) {
	for _, tt := range reduceInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceInt16(tt.input, func(acc int16, elem int16) int16 {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceInt32Tests = []struct {
	name     string
	input    []int32
	expected func() int32
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
		expected: func() int32 {
			input := []int32{
				0,
				1,
				2,
				3,
			}
			var accumulator int32
			for _, in := range input {
			  accumulator = int32(accumulator) + int32(in)
			}
			return accumulator
		},
	},
}

func TestReduceInt32(t *testing.T) {
	for _, tt := range reduceInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceInt32(tt.input, func(acc int32, elem int32) int32 {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceInt64Tests = []struct {
	name     string
	input    []int64
	expected func() int64
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
		expected: func() int64 {
			input := []int64{
				0,
				1,
				2,
				3,
			}
			var accumulator int64
			for _, in := range input {
			  accumulator = int64(accumulator) + int64(in)
			}
			return accumulator
		},
	},
}

func TestReduceInt64(t *testing.T) {
	for _, tt := range reduceInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceInt64(tt.input, func(acc int64, elem int64) int64 {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceUintTests = []struct {
	name     string
	input    []uint
	expected func() uint
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
		expected: func() uint {
			input := []uint{
				0,
				1,
				2,
				3,
			}
			var accumulator uint
			for _, in := range input {
			  accumulator = uint(accumulator) + uint(in)
			}
			return accumulator
		},
	},
}

func TestReduceUint(t *testing.T) {
	for _, tt := range reduceUintTests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceUint(tt.input, func(acc uint, elem uint) uint {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceUint8Tests = []struct {
	name     string
	input    []uint8
	expected func() uint8
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
		expected: func() uint8 {
			input := []uint8{
				0,
				1,
				2,
				3,
			}
			var accumulator uint8
			for _, in := range input {
			  accumulator = uint8(accumulator) + uint8(in)
			}
			return accumulator
		},
	},
}

func TestReduceUint8(t *testing.T) {
	for _, tt := range reduceUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceUint8(tt.input, func(acc uint8, elem uint8) uint8 {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceUint16Tests = []struct {
	name     string
	input    []uint16
	expected func() uint16
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
		expected: func() uint16 {
			input := []uint16{
				0,
				1,
				2,
				3,
			}
			var accumulator uint16
			for _, in := range input {
			  accumulator = uint16(accumulator) + uint16(in)
			}
			return accumulator
		},
	},
}

func TestReduceUint16(t *testing.T) {
	for _, tt := range reduceUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceUint16(tt.input, func(acc uint16, elem uint16) uint16 {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceUint32Tests = []struct {
	name     string
	input    []uint32
	expected func() uint32
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
		expected: func() uint32 {
			input := []uint32{
				0,
				1,
				2,
				3,
			}
			var accumulator uint32
			for _, in := range input {
			  accumulator = uint32(accumulator) + uint32(in)
			}
			return accumulator
		},
	},
}

func TestReduceUint32(t *testing.T) {
	for _, tt := range reduceUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceUint32(tt.input, func(acc uint32, elem uint32) uint32 {
                return acc + elem
            }, 0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceFloat32Tests = []struct {
	name     string
	input    []float32
	expected func() float32
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
		expected: func() float32 {
			input := []float32{
				0.0,
				1.1,
				2.2,
				3.3,
			}
			var accumulator float32
			for _, in := range input {
			  accumulator = float32(accumulator) + float32(in)
			}
			return accumulator
		},
	},
}

func TestReduceFloat32(t *testing.T) {
	for _, tt := range reduceFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceFloat32(tt.input, func(acc float32, elem float32) float32 {
                return acc + elem
            }, 0.0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceFloat64Tests = []struct {
	name     string
	input    []float64
	expected func() float64
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
		expected: func() float64 {
			input := []float64{
				0.0,
				1.1,
				2.2,
				3.3,
			}
			var accumulator float64
			for _, in := range input {
			  accumulator = float64(accumulator) + float64(in)
			}
			return accumulator
		},
	},
}

func TestReduceFloat64(t *testing.T) {
	for _, tt := range reduceFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceFloat64(tt.input, func(acc float64, elem float64) float64 {
                return acc + elem
            }, 0.0)
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceComplex64Tests = []struct {
	name     string
	input    []complex64
	expected func() complex64
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
		expected: func() complex64 {
			input := []complex64{
				(0+3.14i),
				(2.4+3.14i),
				(4.8+3.14i),
				(7.2+3.14i),
			}
			var accumulator complex64
			for _, in := range input {
			  accumulator = complex64(accumulator) + complex64(in)
			}
			return accumulator
		},
	},
}

func TestReduceComplex64(t *testing.T) {
	for _, tt := range reduceComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceComplex64(tt.input, func(acc complex64, elem complex64) complex64 {
                return acc + elem
            }, (0+3.14i))
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}

var reduceComplex128Tests = []struct {
	name     string
	input    []complex128
	expected func() complex128
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
		expected: func() complex128 {
			input := []complex128{
				(0+3.14i),
				(2.4+3.14i),
				(4.8+3.14i),
				(7.199999999999999+3.14i),
			}
			var accumulator complex128
			for _, in := range input {
			  accumulator = complex128(accumulator) + complex128(in)
			}
			return accumulator
		},
	},
}

func TestReduceComplex128(t *testing.T) {
	for _, tt := range reduceComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := ReduceComplex128(tt.input, func(acc complex128, elem complex128) complex128 {
                return acc + elem
            }, (0+3.14i))
			expectedValue := tt.expected()
			assert.Equal(t, output, expectedValue)
		})
	}

}
