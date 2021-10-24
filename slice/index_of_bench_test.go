package slice

import (
	"testing"
)
		

var indexOfBenchStringTests = []struct {
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

func BenchmarkTestIndexOfString(b *testing.B) {
	for _, tt := range indexOfBenchStringTests {
	    for i := 0; i < b.N; i++ {
			IndexOfString(tt.element, tt.input)
		}
	}
}


var indexOfBenchIntTests = []struct {
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

func BenchmarkTestIndexOfInt(b *testing.B) {
	for _, tt := range indexOfBenchIntTests {
	    for i := 0; i < b.N; i++ {
			IndexOfInt(tt.element, tt.input)
		}
	}
}


var indexOfBenchInt8Tests = []struct {
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

func BenchmarkTestIndexOfInt8(b *testing.B) {
	for _, tt := range indexOfBenchInt8Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfInt8(tt.element, tt.input)
		}
	}
}


var indexOfBenchInt16Tests = []struct {
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

func BenchmarkTestIndexOfInt16(b *testing.B) {
	for _, tt := range indexOfBenchInt16Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfInt16(tt.element, tt.input)
		}
	}
}


var indexOfBenchInt32Tests = []struct {
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

func BenchmarkTestIndexOfInt32(b *testing.B) {
	for _, tt := range indexOfBenchInt32Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfInt32(tt.element, tt.input)
		}
	}
}


var indexOfBenchInt64Tests = []struct {
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

func BenchmarkTestIndexOfInt64(b *testing.B) {
	for _, tt := range indexOfBenchInt64Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfInt64(tt.element, tt.input)
		}
	}
}


var indexOfBenchUintTests = []struct {
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

func BenchmarkTestIndexOfUint(b *testing.B) {
	for _, tt := range indexOfBenchUintTests {
	    for i := 0; i < b.N; i++ {
			IndexOfUint(tt.element, tt.input)
		}
	}
}


var indexOfBenchUint8Tests = []struct {
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

func BenchmarkTestIndexOfUint8(b *testing.B) {
	for _, tt := range indexOfBenchUint8Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfUint8(tt.element, tt.input)
		}
	}
}


var indexOfBenchUint16Tests = []struct {
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

func BenchmarkTestIndexOfUint16(b *testing.B) {
	for _, tt := range indexOfBenchUint16Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfUint16(tt.element, tt.input)
		}
	}
}


var indexOfBenchUint32Tests = []struct {
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

func BenchmarkTestIndexOfUint32(b *testing.B) {
	for _, tt := range indexOfBenchUint32Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfUint32(tt.element, tt.input)
		}
	}
}


var indexOfBenchFloat32Tests = []struct {
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

func BenchmarkTestIndexOfFloat32(b *testing.B) {
	for _, tt := range indexOfBenchFloat32Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfFloat32(tt.element, tt.input)
		}
	}
}


var indexOfBenchFloat64Tests = []struct {
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

func BenchmarkTestIndexOfFloat64(b *testing.B) {
	for _, tt := range indexOfBenchFloat64Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfFloat64(tt.element, tt.input)
		}
	}
}


var indexOfBenchComplex64Tests = []struct {
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

func BenchmarkTestIndexOfComplex64(b *testing.B) {
	for _, tt := range indexOfBenchComplex64Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfComplex64(tt.element, tt.input)
		}
	}
}


var indexOfBenchComplex128Tests = []struct {
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

func BenchmarkTestIndexOfComplex128(b *testing.B) {
	for _, tt := range indexOfBenchComplex128Tests {
	    for i := 0; i < b.N; i++ {
			IndexOfComplex128(tt.element, tt.input)
		}
	}
}
