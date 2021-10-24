package slice

import (
	"testing"
)
		
var flattenBenchStringTests = []struct {
	name     string
	input1   []string
	input2   []string
	expected []string
}{

	{
		name: "basic example",
		input1: []string{
			"test1",
			"test2",
			"test3",
		},
		input2: []string{
			"test1",
			"test2",
			"test3",
		},
		expected: []string{
            "test1",
            "test2",
            "test3",
            "test1",
            "test2",
            "test3",
        },
	},
}

func BenchmarkTestFlattenString(b *testing.B) {
	for _, tt := range flattenBenchStringTests {
	    for i := 0; i < b.N; i++ {
			FlattenString(tt.input1, tt.input2)
		}
	}

}

var flattenBenchIntTests = []struct {
	name     string
	input1   []int
	input2   []int
	expected []int
}{

	{
		name: "basic example",
		input1: []int{
			1,
			2,
			3,
		},
		input2: []int{
			1,
			2,
			3,
		},
		expected: []int{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenInt(b *testing.B) {
	for _, tt := range flattenBenchIntTests {
	    for i := 0; i < b.N; i++ {
			FlattenInt(tt.input1, tt.input2)
		}
	}

}

var flattenBenchInt8Tests = []struct {
	name     string
	input1   []int8
	input2   []int8
	expected []int8
}{

	{
		name: "basic example",
		input1: []int8{
			1,
			2,
			3,
		},
		input2: []int8{
			1,
			2,
			3,
		},
		expected: []int8{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenInt8(b *testing.B) {
	for _, tt := range flattenBenchInt8Tests {
	    for i := 0; i < b.N; i++ {
			FlattenInt8(tt.input1, tt.input2)
		}
	}

}

var flattenBenchInt16Tests = []struct {
	name     string
	input1   []int16
	input2   []int16
	expected []int16
}{

	{
		name: "basic example",
		input1: []int16{
			1,
			2,
			3,
		},
		input2: []int16{
			1,
			2,
			3,
		},
		expected: []int16{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenInt16(b *testing.B) {
	for _, tt := range flattenBenchInt16Tests {
	    for i := 0; i < b.N; i++ {
			FlattenInt16(tt.input1, tt.input2)
		}
	}

}

var flattenBenchInt32Tests = []struct {
	name     string
	input1   []int32
	input2   []int32
	expected []int32
}{

	{
		name: "basic example",
		input1: []int32{
			1,
			2,
			3,
		},
		input2: []int32{
			1,
			2,
			3,
		},
		expected: []int32{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenInt32(b *testing.B) {
	for _, tt := range flattenBenchInt32Tests {
	    for i := 0; i < b.N; i++ {
			FlattenInt32(tt.input1, tt.input2)
		}
	}

}

var flattenBenchInt64Tests = []struct {
	name     string
	input1   []int64
	input2   []int64
	expected []int64
}{

	{
		name: "basic example",
		input1: []int64{
			1,
			2,
			3,
		},
		input2: []int64{
			1,
			2,
			3,
		},
		expected: []int64{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenInt64(b *testing.B) {
	for _, tt := range flattenBenchInt64Tests {
	    for i := 0; i < b.N; i++ {
			FlattenInt64(tt.input1, tt.input2)
		}
	}

}

var flattenBenchUintTests = []struct {
	name     string
	input1   []uint
	input2   []uint
	expected []uint
}{

	{
		name: "basic example",
		input1: []uint{
			1,
			2,
			3,
		},
		input2: []uint{
			1,
			2,
			3,
		},
		expected: []uint{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenUint(b *testing.B) {
	for _, tt := range flattenBenchUintTests {
	    for i := 0; i < b.N; i++ {
			FlattenUint(tt.input1, tt.input2)
		}
	}

}

var flattenBenchUint8Tests = []struct {
	name     string
	input1   []uint8
	input2   []uint8
	expected []uint8
}{

	{
		name: "basic example",
		input1: []uint8{
			1,
			2,
			3,
		},
		input2: []uint8{
			1,
			2,
			3,
		},
		expected: []uint8{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenUint8(b *testing.B) {
	for _, tt := range flattenBenchUint8Tests {
	    for i := 0; i < b.N; i++ {
			FlattenUint8(tt.input1, tt.input2)
		}
	}

}

var flattenBenchUint16Tests = []struct {
	name     string
	input1   []uint16
	input2   []uint16
	expected []uint16
}{

	{
		name: "basic example",
		input1: []uint16{
			1,
			2,
			3,
		},
		input2: []uint16{
			1,
			2,
			3,
		},
		expected: []uint16{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenUint16(b *testing.B) {
	for _, tt := range flattenBenchUint16Tests {
	    for i := 0; i < b.N; i++ {
			FlattenUint16(tt.input1, tt.input2)
		}
	}

}

var flattenBenchUint32Tests = []struct {
	name     string
	input1   []uint32
	input2   []uint32
	expected []uint32
}{

	{
		name: "basic example",
		input1: []uint32{
			1,
			2,
			3,
		},
		input2: []uint32{
			1,
			2,
			3,
		},
		expected: []uint32{
            1,
            2,
            3,
            1,
            2,
            3,
        },
	},
}

func BenchmarkTestFlattenUint32(b *testing.B) {
	for _, tt := range flattenBenchUint32Tests {
	    for i := 0; i < b.N; i++ {
			FlattenUint32(tt.input1, tt.input2)
		}
	}

}

var flattenBenchFloat32Tests = []struct {
	name     string
	input1   []float32
	input2   []float32
	expected []float32
}{

	{
		name: "basic example",
		input1: []float32{
			1.1,
			2.2,
			3.3,
		},
		input2: []float32{
			1.1,
			2.2,
			3.3,
		},
		expected: []float32{
            1.1,
            2.2,
            3.3,
            1.1,
            2.2,
            3.3,
        },
	},
}

func BenchmarkTestFlattenFloat32(b *testing.B) {
	for _, tt := range flattenBenchFloat32Tests {
	    for i := 0; i < b.N; i++ {
			FlattenFloat32(tt.input1, tt.input2)
		}
	}

}

var flattenBenchFloat64Tests = []struct {
	name     string
	input1   []float64
	input2   []float64
	expected []float64
}{

	{
		name: "basic example",
		input1: []float64{
			1.1,
			2.2,
			3.3,
		},
		input2: []float64{
			1.1,
			2.2,
			3.3,
		},
		expected: []float64{
            1.1,
            2.2,
            3.3,
            1.1,
            2.2,
            3.3,
        },
	},
}

func BenchmarkTestFlattenFloat64(b *testing.B) {
	for _, tt := range flattenBenchFloat64Tests {
	    for i := 0; i < b.N; i++ {
			FlattenFloat64(tt.input1, tt.input2)
		}
	}

}

var flattenBenchComplex64Tests = []struct {
	name     string
	input1   []complex64
	input2   []complex64
	expected []complex64
}{

	{
		name: "basic example",
		input1: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
		input2: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
		expected: []complex64{
            (2.4+3.14i),
            (4.8+3.14i),
            (7.2+3.14i),
            (2.4+3.14i),
            (4.8+3.14i),
            (7.2+3.14i),
        },
	},
}

func BenchmarkTestFlattenComplex64(b *testing.B) {
	for _, tt := range flattenBenchComplex64Tests {
	    for i := 0; i < b.N; i++ {
			FlattenComplex64(tt.input1, tt.input2)
		}
	}

}

var flattenBenchComplex128Tests = []struct {
	name     string
	input1   []complex128
	input2   []complex128
	expected []complex128
}{

	{
		name: "basic example",
		input1: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
		input2: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
		expected: []complex128{
            (2.4+3.14i),
            (4.8+3.14i),
            (7.199999999999999+3.14i),
            (2.4+3.14i),
            (4.8+3.14i),
            (7.199999999999999+3.14i),
        },
	},
}

func BenchmarkTestFlattenComplex128(b *testing.B) {
	for _, tt := range flattenBenchComplex128Tests {
	    for i := 0; i < b.N; i++ {
			FlattenComplex128(tt.input1, tt.input2)
		}
	}

}
