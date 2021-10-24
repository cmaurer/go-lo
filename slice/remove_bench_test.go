package slice

import (
	"testing"
)
		
var removeBenchStringTests = []struct {
	name     	string
	input    	[]string
	equalTo  	string
	expected 	[]string
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
        equalTo: "test2",
		expected: []string{ "test1", "test3" },
	},
	{
		name: "remove all",
		input: []string{
			"test1",
		},
        equalTo: "test1",
		expected: []string{},
	},
	{
		name: "remove none",
		input: []string{
			"test1",
			"test2",
		},
        equalTo: "test3",
		expected: []string{ "test1", "test2" },
	},

}

func BenchmarkTestRemoveString(b *testing.B) {
	for _, tt := range removeBenchStringTests {
		for i := 0; i < b.N; i++ {
			RemoveString(tt.input, func(in string) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchIntTests = []struct {
	name     	string
	input    	[]int
	equalTo  	int
	expected 	[]int
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int{ 1, 3 },
	},
	{
		name: "remove all",
		input: []int{
			1,
		},
        equalTo: 1,
		expected: []int{},
	},
	{
		name: "remove none",
		input: []int{
			1,
			2,
		},
        equalTo: 3,
		expected: []int{ 1, 2 },
	},

}

func BenchmarkTestRemoveInt(b *testing.B) {
	for _, tt := range removeBenchIntTests {
		for i := 0; i < b.N; i++ {
			RemoveInt(tt.input, func(in int) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchInt8Tests = []struct {
	name     	string
	input    	[]int8
	equalTo  	int8
	expected 	[]int8
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int8{ 1, 3 },
	},
	{
		name: "remove all",
		input: []int8{
			1,
		},
        equalTo: 1,
		expected: []int8{},
	},
	{
		name: "remove none",
		input: []int8{
			1,
			2,
		},
        equalTo: 3,
		expected: []int8{ 1, 2 },
	},

}

func BenchmarkTestRemoveInt8(b *testing.B) {
	for _, tt := range removeBenchInt8Tests {
		for i := 0; i < b.N; i++ {
			RemoveInt8(tt.input, func(in int8) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchInt16Tests = []struct {
	name     	string
	input    	[]int16
	equalTo  	int16
	expected 	[]int16
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int16{ 1, 3 },
	},
	{
		name: "remove all",
		input: []int16{
			1,
		},
        equalTo: 1,
		expected: []int16{},
	},
	{
		name: "remove none",
		input: []int16{
			1,
			2,
		},
        equalTo: 3,
		expected: []int16{ 1, 2 },
	},

}

func BenchmarkTestRemoveInt16(b *testing.B) {
	for _, tt := range removeBenchInt16Tests {
		for i := 0; i < b.N; i++ {
			RemoveInt16(tt.input, func(in int16) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchInt32Tests = []struct {
	name     	string
	input    	[]int32
	equalTo  	int32
	expected 	[]int32
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int32{ 1, 3 },
	},
	{
		name: "remove all",
		input: []int32{
			1,
		},
        equalTo: 1,
		expected: []int32{},
	},
	{
		name: "remove none",
		input: []int32{
			1,
			2,
		},
        equalTo: 3,
		expected: []int32{ 1, 2 },
	},

}

func BenchmarkTestRemoveInt32(b *testing.B) {
	for _, tt := range removeBenchInt32Tests {
		for i := 0; i < b.N; i++ {
			RemoveInt32(tt.input, func(in int32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchInt64Tests = []struct {
	name     	string
	input    	[]int64
	equalTo  	int64
	expected 	[]int64
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int64{ 1, 3 },
	},
	{
		name: "remove all",
		input: []int64{
			1,
		},
        equalTo: 1,
		expected: []int64{},
	},
	{
		name: "remove none",
		input: []int64{
			1,
			2,
		},
        equalTo: 3,
		expected: []int64{ 1, 2 },
	},

}

func BenchmarkTestRemoveInt64(b *testing.B) {
	for _, tt := range removeBenchInt64Tests {
		for i := 0; i < b.N; i++ {
			RemoveInt64(tt.input, func(in int64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchUintTests = []struct {
	name     	string
	input    	[]uint
	equalTo  	uint
	expected 	[]uint
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint{ 1, 3 },
	},
	{
		name: "remove all",
		input: []uint{
			1,
		},
        equalTo: 1,
		expected: []uint{},
	},
	{
		name: "remove none",
		input: []uint{
			1,
			2,
		},
        equalTo: 3,
		expected: []uint{ 1, 2 },
	},

}

func BenchmarkTestRemoveUint(b *testing.B) {
	for _, tt := range removeBenchUintTests {
		for i := 0; i < b.N; i++ {
			RemoveUint(tt.input, func(in uint) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchUint8Tests = []struct {
	name     	string
	input    	[]uint8
	equalTo  	uint8
	expected 	[]uint8
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint8{ 1, 3 },
	},
	{
		name: "remove all",
		input: []uint8{
			1,
		},
        equalTo: 1,
		expected: []uint8{},
	},
	{
		name: "remove none",
		input: []uint8{
			1,
			2,
		},
        equalTo: 3,
		expected: []uint8{ 1, 2 },
	},

}

func BenchmarkTestRemoveUint8(b *testing.B) {
	for _, tt := range removeBenchUint8Tests {
		for i := 0; i < b.N; i++ {
			RemoveUint8(tt.input, func(in uint8) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchUint16Tests = []struct {
	name     	string
	input    	[]uint16
	equalTo  	uint16
	expected 	[]uint16
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint16{ 1, 3 },
	},
	{
		name: "remove all",
		input: []uint16{
			1,
		},
        equalTo: 1,
		expected: []uint16{},
	},
	{
		name: "remove none",
		input: []uint16{
			1,
			2,
		},
        equalTo: 3,
		expected: []uint16{ 1, 2 },
	},

}

func BenchmarkTestRemoveUint16(b *testing.B) {
	for _, tt := range removeBenchUint16Tests {
		for i := 0; i < b.N; i++ {
			RemoveUint16(tt.input, func(in uint16) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchUint32Tests = []struct {
	name     	string
	input    	[]uint32
	equalTo  	uint32
	expected 	[]uint32
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint32{ 1, 3 },
	},
	{
		name: "remove all",
		input: []uint32{
			1,
		},
        equalTo: 1,
		expected: []uint32{},
	},
	{
		name: "remove none",
		input: []uint32{
			1,
			2,
		},
        equalTo: 3,
		expected: []uint32{ 1, 2 },
	},

}

func BenchmarkTestRemoveUint32(b *testing.B) {
	for _, tt := range removeBenchUint32Tests {
		for i := 0; i < b.N; i++ {
			RemoveUint32(tt.input, func(in uint32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchFloat32Tests = []struct {
	name     	string
	input    	[]float32
	equalTo  	float32
	expected 	[]float32
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
        equalTo: 2.2,
		expected: []float32{ 1.1, 3.3 },
	},
	{
		name: "remove all",
		input: []float32{
			1.1,
		},
        equalTo: 1.1,
		expected: []float32{},
	},
	{
		name: "remove none",
		input: []float32{
			1.1,
			2.2,
		},
        equalTo: 3.3,
		expected: []float32{ 1.1, 2.2 },
	},

}

func BenchmarkTestRemoveFloat32(b *testing.B) {
	for _, tt := range removeBenchFloat32Tests {
		for i := 0; i < b.N; i++ {
			RemoveFloat32(tt.input, func(in float32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchFloat64Tests = []struct {
	name     	string
	input    	[]float64
	equalTo  	float64
	expected 	[]float64
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
        equalTo: 2.2,
		expected: []float64{ 1.1, 3.3 },
	},
	{
		name: "remove all",
		input: []float64{
			1.1,
		},
        equalTo: 1.1,
		expected: []float64{},
	},
	{
		name: "remove none",
		input: []float64{
			1.1,
			2.2,
		},
        equalTo: 3.3,
		expected: []float64{ 1.1, 2.2 },
	},

}

func BenchmarkTestRemoveFloat64(b *testing.B) {
	for _, tt := range removeBenchFloat64Tests {
		for i := 0; i < b.N; i++ {
			RemoveFloat64(tt.input, func(in float64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchComplex64Tests = []struct {
	name     	string
	input    	[]complex64
	equalTo  	complex64
	expected 	[]complex64
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: []complex64{ (2.4+3.14i), (7.2+3.14i) },
	},
	{
		name: "remove all",
		input: []complex64{
			(2.4+3.14i),
		},
        equalTo: (2.4+3.14i),
		expected: []complex64{},
	},
	{
		name: "remove none",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
        equalTo: (7.2+3.14i),
		expected: []complex64{ (2.4+3.14i), (4.8+3.14i) },
	},

}

func BenchmarkTestRemoveComplex64(b *testing.B) {
	for _, tt := range removeBenchComplex64Tests {
		for i := 0; i < b.N; i++ {
			RemoveComplex64(tt.input, func(in complex64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var removeBenchComplex128Tests = []struct {
	name     	string
	input    	[]complex128
	equalTo  	complex128
	expected 	[]complex128
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: []complex128{ (2.4+3.14i), (7.199999999999999+3.14i) },
	},
	{
		name: "remove all",
		input: []complex128{
			(2.4+3.14i),
		},
        equalTo: (2.4+3.14i),
		expected: []complex128{},
	},
	{
		name: "remove none",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
        equalTo: (7.199999999999999+3.14i),
		expected: []complex128{ (2.4+3.14i), (4.8+3.14i) },
	},

}

func BenchmarkTestRemoveComplex128(b *testing.B) {
	for _, tt := range removeBenchComplex128Tests {
		for i := 0; i < b.N; i++ {
			RemoveComplex128(tt.input, func(in complex128) bool {
				return in == tt.equalTo
			})
		}
	}

}
