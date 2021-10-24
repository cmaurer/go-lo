package slice

import (
	"testing"
)
		


var mapBenchStringTests = []struct {
	name     string
	input    []string
	expected []string
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
		expected: []string{
            "test1"+"test1",
            "test2"+"test2",
            "test3"+"test3",
        },
	},
}

func BenchmarkTestMapString(b *testing.B) {
	for _, tt := range mapBenchStringTests {
		MapString(tt.input, func(in string) string {
			return in + in
		})
	}

}



var mapBenchIntTests = []struct {
	name     string
	input    []int
	expected []int
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
		expected: []int{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapInt(b *testing.B) {
	for _, tt := range mapBenchIntTests {
		MapInt(tt.input, func(in int) int {
			return in + in
		})
	}

}



var mapBenchInt8Tests = []struct {
	name     string
	input    []int8
	expected []int8
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
		expected: []int8{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapInt8(b *testing.B) {
	for _, tt := range mapBenchInt8Tests {
		MapInt8(tt.input, func(in int8) int8 {
			return in + in
		})
	}

}



var mapBenchInt16Tests = []struct {
	name     string
	input    []int16
	expected []int16
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
		expected: []int16{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapInt16(b *testing.B) {
	for _, tt := range mapBenchInt16Tests {
		MapInt16(tt.input, func(in int16) int16 {
			return in + in
		})
	}

}



var mapBenchInt32Tests = []struct {
	name     string
	input    []int32
	expected []int32
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
		expected: []int32{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapInt32(b *testing.B) {
	for _, tt := range mapBenchInt32Tests {
		MapInt32(tt.input, func(in int32) int32 {
			return in + in
		})
	}

}



var mapBenchInt64Tests = []struct {
	name     string
	input    []int64
	expected []int64
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
		expected: []int64{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapInt64(b *testing.B) {
	for _, tt := range mapBenchInt64Tests {
		MapInt64(tt.input, func(in int64) int64 {
			return in + in
		})
	}

}



var mapBenchUintTests = []struct {
	name     string
	input    []uint
	expected []uint
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
		expected: []uint{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapUint(b *testing.B) {
	for _, tt := range mapBenchUintTests {
		MapUint(tt.input, func(in uint) uint {
			return in + in
		})
	}

}



var mapBenchUint8Tests = []struct {
	name     string
	input    []uint8
	expected []uint8
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
		expected: []uint8{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapUint8(b *testing.B) {
	for _, tt := range mapBenchUint8Tests {
		MapUint8(tt.input, func(in uint8) uint8 {
			return in + in
		})
	}

}



var mapBenchUint16Tests = []struct {
	name     string
	input    []uint16
	expected []uint16
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
		expected: []uint16{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapUint16(b *testing.B) {
	for _, tt := range mapBenchUint16Tests {
		MapUint16(tt.input, func(in uint16) uint16 {
			return in + in
		})
	}

}



var mapBenchUint32Tests = []struct {
	name     string
	input    []uint32
	expected []uint32
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
		expected: []uint32{
            1+1,
            2+2,
            3+3,
        },
	},
}

func BenchmarkTestMapUint32(b *testing.B) {
	for _, tt := range mapBenchUint32Tests {
		MapUint32(tt.input, func(in uint32) uint32 {
			return in + in
		})
	}

}



var mapBenchFloat32Tests = []struct {
	name     string
	input    []float32
	expected []float32
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
		expected: []float32{
            1.1+1.1,
            2.2+2.2,
            3.3+3.3,
        },
	},
}

func BenchmarkTestMapFloat32(b *testing.B) {
	for _, tt := range mapBenchFloat32Tests {
		MapFloat32(tt.input, func(in float32) float32 {
			return in + in
		})
	}

}



var mapBenchFloat64Tests = []struct {
	name     string
	input    []float64
	expected []float64
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
		expected: []float64{
            1.1+1.1,
            2.2+2.2,
            3.3+3.3,
        },
	},
}

func BenchmarkTestMapFloat64(b *testing.B) {
	for _, tt := range mapBenchFloat64Tests {
		MapFloat64(tt.input, func(in float64) float64 {
			return in + in
		})
	}

}



var mapBenchComplex64Tests = []struct {
	name     string
	input    []complex64
	expected []complex64
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
		expected: []complex64{
            (2.4+3.14i)+(2.4+3.14i),
            (4.8+3.14i)+(4.8+3.14i),
            (7.2+3.14i)+(7.2+3.14i),
        },
	},
}

func BenchmarkTestMapComplex64(b *testing.B) {
	for _, tt := range mapBenchComplex64Tests {
		MapComplex64(tt.input, func(in complex64) complex64 {
			return in + in
		})
	}

}



var mapBenchComplex128Tests = []struct {
	name     string
	input    []complex128
	expected []complex128
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
		expected: []complex128{
            (2.4+3.14i)+(2.4+3.14i),
            (4.8+3.14i)+(4.8+3.14i),
            (7.199999999999999+3.14i)+(7.199999999999999+3.14i),
        },
	},
}

func BenchmarkTestMapComplex128(b *testing.B) {
	for _, tt := range mapBenchComplex128Tests {
		MapComplex128(tt.input, func(in complex128) complex128 {
			return in + in
		})
	}

}
