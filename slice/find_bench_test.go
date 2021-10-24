package slice

import (
	"testing"
)
		
var findBenchStringTests = []struct {
	name     	string
	input    	[]string
	equalTo  	string
	expected 	string
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
        equalTo: "test2",
		expected: "test2",
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []string{
			"test1",
			"test2",
			"test3",
			"test2",
		},
        equalTo: "test2",
		expected: "test2",
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []string{
			"test1",
			"test2",
			"test3",
			"test2",
		},
        equalTo: "test8",
		expected: "test8",
		shouldBeNil: true,
	},
}

func BenchmarkTestFindString(b *testing.B) {
	for _, tt := range findStringTests {
    	for i := 0; i < b.N; i++ {
			FindString(tt.input, func(in string) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchIntTests = []struct {
	name     	string
	input    	[]int
	equalTo  	int
	expected 	int
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []int{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []int{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindInt(b *testing.B) {
	for _, tt := range findIntTests {
    	for i := 0; i < b.N; i++ {
			FindInt(tt.input, func(in int) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchInt8Tests = []struct {
	name     	string
	input    	[]int8
	equalTo  	int8
	expected 	int8
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []int8{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []int8{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindInt8(b *testing.B) {
	for _, tt := range findInt8Tests {
    	for i := 0; i < b.N; i++ {
			FindInt8(tt.input, func(in int8) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchInt16Tests = []struct {
	name     	string
	input    	[]int16
	equalTo  	int16
	expected 	int16
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []int16{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []int16{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindInt16(b *testing.B) {
	for _, tt := range findInt16Tests {
    	for i := 0; i < b.N; i++ {
			FindInt16(tt.input, func(in int16) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchInt32Tests = []struct {
	name     	string
	input    	[]int32
	equalTo  	int32
	expected 	int32
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []int32{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []int32{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindInt32(b *testing.B) {
	for _, tt := range findInt32Tests {
    	for i := 0; i < b.N; i++ {
			FindInt32(tt.input, func(in int32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchInt64Tests = []struct {
	name     	string
	input    	[]int64
	equalTo  	int64
	expected 	int64
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []int64{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []int64{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindInt64(b *testing.B) {
	for _, tt := range findInt64Tests {
    	for i := 0; i < b.N; i++ {
			FindInt64(tt.input, func(in int64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchUintTests = []struct {
	name     	string
	input    	[]uint
	equalTo  	uint
	expected 	uint
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []uint{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []uint{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindUint(b *testing.B) {
	for _, tt := range findUintTests {
    	for i := 0; i < b.N; i++ {
			FindUint(tt.input, func(in uint) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchUint8Tests = []struct {
	name     	string
	input    	[]uint8
	equalTo  	uint8
	expected 	uint8
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []uint8{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []uint8{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindUint8(b *testing.B) {
	for _, tt := range findUint8Tests {
    	for i := 0; i < b.N; i++ {
			FindUint8(tt.input, func(in uint8) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchUint16Tests = []struct {
	name     	string
	input    	[]uint16
	equalTo  	uint16
	expected 	uint16
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []uint16{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []uint16{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindUint16(b *testing.B) {
	for _, tt := range findUint16Tests {
    	for i := 0; i < b.N; i++ {
			FindUint16(tt.input, func(in uint16) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchUint32Tests = []struct {
	name     	string
	input    	[]uint32
	equalTo  	uint32
	expected 	uint32
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []uint32{
			1,
			2,
			3,
			2,
		},
        equalTo: 2,
		expected: 2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []uint32{
			1,
			2,
			3,
			2,
		},
        equalTo: 8,
		expected: 8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindUint32(b *testing.B) {
	for _, tt := range findUint32Tests {
    	for i := 0; i < b.N; i++ {
			FindUint32(tt.input, func(in uint32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchFloat32Tests = []struct {
	name     	string
	input    	[]float32
	equalTo  	float32
	expected 	float32
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
        equalTo: 2.2,
		expected: 2.2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []float32{
			1.1,
			2.2,
			3.3,
			2.2,
		},
        equalTo: 2.2,
		expected: 2.2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []float32{
			1.1,
			2.2,
			3.3,
			2.2,
		},
        equalTo: 8.8,
		expected: 8.8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindFloat32(b *testing.B) {
	for _, tt := range findFloat32Tests {
    	for i := 0; i < b.N; i++ {
			FindFloat32(tt.input, func(in float32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchFloat64Tests = []struct {
	name     	string
	input    	[]float64
	equalTo  	float64
	expected 	float64
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
        equalTo: 2.2,
		expected: 2.2,
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []float64{
			1.1,
			2.2,
			3.3,
			2.2,
		},
        equalTo: 2.2,
		expected: 2.2,
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []float64{
			1.1,
			2.2,
			3.3,
			2.2,
		},
        equalTo: 8.8,
		expected: 8.8,
		shouldBeNil: true,
	},
}

func BenchmarkTestFindFloat64(b *testing.B) {
	for _, tt := range findFloat64Tests {
    	for i := 0; i < b.N; i++ {
			FindFloat64(tt.input, func(in float64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchComplex64Tests = []struct {
	name     	string
	input    	[]complex64
	equalTo  	complex64
	expected 	complex64
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: (4.8+3.14i),
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
			(4.8+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: (4.8+3.14i),
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
			(4.8+3.14i),
		},
        equalTo: (19.2+3.14i),
		expected: (19.2+3.14i),
		shouldBeNil: true,
	},
}

func BenchmarkTestFindComplex64(b *testing.B) {
	for _, tt := range findComplex64Tests {
    	for i := 0; i < b.N; i++ {
			FindComplex64(tt.input, func(in complex64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var findBenchComplex128Tests = []struct {
	name     	string
	input    	[]complex128
	equalTo  	complex128
	expected 	complex128
	shouldBeNil bool
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: (4.8+3.14i),
		shouldBeNil: false,
	},
	{
		name: "more than one element",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
			(4.8+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: (4.8+3.14i),
		shouldBeNil: false,
	},
	{
		name: "not found",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
			(4.8+3.14i),
		},
        equalTo: (19.2+3.14i),
		expected: (19.2+3.14i),
		shouldBeNil: true,
	},
}

func BenchmarkTestFindComplex128(b *testing.B) {
	for _, tt := range findComplex128Tests {
    	for i := 0; i < b.N; i++ {
			FindComplex128(tt.input, func(in complex128) bool {
				return in == tt.equalTo
			})
		}
	}

}
