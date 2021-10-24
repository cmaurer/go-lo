package slice

import (
	"testing"
)
		
var includesBenchStringTests = []struct {
	name     	string
	input    	[]string
	included  	string
	expected 	bool
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

func BenchmarkTestIncludesString(b *testing.B) {
	for _, tt := range includesBenchStringTests {
		for i := 0; i < b.N; i++ {
			IncludesString(tt.input, tt.included)
		}
	}

}

var includesBenchIntTests = []struct {
	name     	string
	input    	[]int
	included  	int
	expected 	bool
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

func BenchmarkTestIncludesInt(b *testing.B) {
	for _, tt := range includesBenchIntTests {
		for i := 0; i < b.N; i++ {
			IncludesInt(tt.input, tt.included)
		}
	}

}

var includesBenchInt8Tests = []struct {
	name     	string
	input    	[]int8
	included  	int8
	expected 	bool
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

func BenchmarkTestIncludesInt8(b *testing.B) {
	for _, tt := range includesBenchInt8Tests {
		for i := 0; i < b.N; i++ {
			IncludesInt8(tt.input, tt.included)
		}
	}

}

var includesBenchInt16Tests = []struct {
	name     	string
	input    	[]int16
	included  	int16
	expected 	bool
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

func BenchmarkTestIncludesInt16(b *testing.B) {
	for _, tt := range includesBenchInt16Tests {
		for i := 0; i < b.N; i++ {
			IncludesInt16(tt.input, tt.included)
		}
	}

}

var includesBenchInt32Tests = []struct {
	name     	string
	input    	[]int32
	included  	int32
	expected 	bool
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

func BenchmarkTestIncludesInt32(b *testing.B) {
	for _, tt := range includesBenchInt32Tests {
		for i := 0; i < b.N; i++ {
			IncludesInt32(tt.input, tt.included)
		}
	}

}

var includesBenchInt64Tests = []struct {
	name     	string
	input    	[]int64
	included  	int64
	expected 	bool
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

func BenchmarkTestIncludesInt64(b *testing.B) {
	for _, tt := range includesBenchInt64Tests {
		for i := 0; i < b.N; i++ {
			IncludesInt64(tt.input, tt.included)
		}
	}

}

var includesBenchUintTests = []struct {
	name     	string
	input    	[]uint
	included  	uint
	expected 	bool
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

func BenchmarkTestIncludesUint(b *testing.B) {
	for _, tt := range includesBenchUintTests {
		for i := 0; i < b.N; i++ {
			IncludesUint(tt.input, tt.included)
		}
	}

}

var includesBenchUint8Tests = []struct {
	name     	string
	input    	[]uint8
	included  	uint8
	expected 	bool
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

func BenchmarkTestIncludesUint8(b *testing.B) {
	for _, tt := range includesBenchUint8Tests {
		for i := 0; i < b.N; i++ {
			IncludesUint8(tt.input, tt.included)
		}
	}

}

var includesBenchUint16Tests = []struct {
	name     	string
	input    	[]uint16
	included  	uint16
	expected 	bool
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

func BenchmarkTestIncludesUint16(b *testing.B) {
	for _, tt := range includesBenchUint16Tests {
		for i := 0; i < b.N; i++ {
			IncludesUint16(tt.input, tt.included)
		}
	}

}

var includesBenchUint32Tests = []struct {
	name     	string
	input    	[]uint32
	included  	uint32
	expected 	bool
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

func BenchmarkTestIncludesUint32(b *testing.B) {
	for _, tt := range includesBenchUint32Tests {
		for i := 0; i < b.N; i++ {
			IncludesUint32(tt.input, tt.included)
		}
	}

}

var includesBenchFloat32Tests = []struct {
	name     	string
	input    	[]float32
	included  	float32
	expected 	bool
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

func BenchmarkTestIncludesFloat32(b *testing.B) {
	for _, tt := range includesBenchFloat32Tests {
		for i := 0; i < b.N; i++ {
			IncludesFloat32(tt.input, tt.included)
		}
	}

}

var includesBenchFloat64Tests = []struct {
	name     	string
	input    	[]float64
	included  	float64
	expected 	bool
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

func BenchmarkTestIncludesFloat64(b *testing.B) {
	for _, tt := range includesBenchFloat64Tests {
		for i := 0; i < b.N; i++ {
			IncludesFloat64(tt.input, tt.included)
		}
	}

}

var includesBenchComplex64Tests = []struct {
	name     	string
	input    	[]complex64
	included  	complex64
	expected 	bool
}{
	{
		name: "should be true",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
        included: (4.8+3.14i),
		expected: true,
	},
	{
		name: "should be false",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
        included: (7.2+3.14i),
		expected: false,
	},
}

func BenchmarkTestIncludesComplex64(b *testing.B) {
	for _, tt := range includesBenchComplex64Tests {
		for i := 0; i < b.N; i++ {
			IncludesComplex64(tt.input, tt.included)
		}
	}

}

var includesBenchComplex128Tests = []struct {
	name     	string
	input    	[]complex128
	included  	complex128
	expected 	bool
}{
	{
		name: "should be true",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
        included: (4.8+3.14i),
		expected: true,
	},
	{
		name: "should be false",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
        included: (7.199999999999999+3.14i),
		expected: false,
	},
}

func BenchmarkTestIncludesComplex128(b *testing.B) {
	for _, tt := range includesBenchComplex128Tests {
		for i := 0; i < b.N; i++ {
			IncludesComplex128(tt.input, tt.included)
		}
	}

}
