package slice

import (
	"testing"
)
		
var eachBenchStringTests = []struct {
	name     string
	input    []string
	expected int
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
        expected: 3,
	},
}

func BenchmarkTestEachString(b *testing.B) {
	for _, tt := range eachBenchStringTests {
		expectedCount := 0
		EachString(tt.input, func(in string) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchIntTests = []struct {
	name     string
	input    []int
	expected int
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachInt(b *testing.B) {
	for _, tt := range eachBenchIntTests {
		expectedCount := 0
		EachInt(tt.input, func(in int) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchInt8Tests = []struct {
	name     string
	input    []int8
	expected int
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachInt8(b *testing.B) {
	for _, tt := range eachBenchInt8Tests {
		expectedCount := 0
		EachInt8(tt.input, func(in int8) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchInt16Tests = []struct {
	name     string
	input    []int16
	expected int
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachInt16(b *testing.B) {
	for _, tt := range eachBenchInt16Tests {
		expectedCount := 0
		EachInt16(tt.input, func(in int16) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchInt32Tests = []struct {
	name     string
	input    []int32
	expected int
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachInt32(b *testing.B) {
	for _, tt := range eachBenchInt32Tests {
		expectedCount := 0
		EachInt32(tt.input, func(in int32) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchInt64Tests = []struct {
	name     string
	input    []int64
	expected int
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachInt64(b *testing.B) {
	for _, tt := range eachBenchInt64Tests {
		expectedCount := 0
		EachInt64(tt.input, func(in int64) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchUintTests = []struct {
	name     string
	input    []uint
	expected int
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachUint(b *testing.B) {
	for _, tt := range eachBenchUintTests {
		expectedCount := 0
		EachUint(tt.input, func(in uint) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchUint8Tests = []struct {
	name     string
	input    []uint8
	expected int
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachUint8(b *testing.B) {
	for _, tt := range eachBenchUint8Tests {
		expectedCount := 0
		EachUint8(tt.input, func(in uint8) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchUint16Tests = []struct {
	name     string
	input    []uint16
	expected int
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachUint16(b *testing.B) {
	for _, tt := range eachBenchUint16Tests {
		expectedCount := 0
		EachUint16(tt.input, func(in uint16) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchUint32Tests = []struct {
	name     string
	input    []uint32
	expected int
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachUint32(b *testing.B) {
	for _, tt := range eachBenchUint32Tests {
		expectedCount := 0
		EachUint32(tt.input, func(in uint32) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchFloat32Tests = []struct {
	name     string
	input    []float32
	expected int
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachFloat32(b *testing.B) {
	for _, tt := range eachBenchFloat32Tests {
		expectedCount := 0
		EachFloat32(tt.input, func(in float32) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchFloat64Tests = []struct {
	name     string
	input    []float64
	expected int
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
        expected: 3,
	},
}

func BenchmarkTestEachFloat64(b *testing.B) {
	for _, tt := range eachBenchFloat64Tests {
		expectedCount := 0
		EachFloat64(tt.input, func(in float64) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchComplex64Tests = []struct {
	name     string
	input    []complex64
	expected int
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
        expected: 3,
	},
}

func BenchmarkTestEachComplex64(b *testing.B) {
	for _, tt := range eachBenchComplex64Tests {
		expectedCount := 0
		EachComplex64(tt.input, func(in complex64) {
			expectedCount = expectedCount + 1
		})
	}

}

var eachBenchComplex128Tests = []struct {
	name     string
	input    []complex128
	expected int
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
        expected: 3,
	},
}

func BenchmarkTestEachComplex128(b *testing.B) {
	for _, tt := range eachBenchComplex128Tests {
		expectedCount := 0
		EachComplex128(tt.input, func(in complex128) {
			expectedCount = expectedCount + 1
		})
	}

}
