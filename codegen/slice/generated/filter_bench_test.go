package slice

import (
	"testing"
)
		
var filterBenchStringTests = []struct {
	name     string
	input    []string
	equalTo  string
	expected []string
}{

	{
		name: "basic example",
		input: []string{
			"test1",
			"test2",
			"test3",
		},
        equalTo: "test2",
		expected: []string{
            "test2",
        },
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
		expected: []string{
            "test2",
            "test2",
        },
	},
}

func BenchmarkTestFilterString(b *testing.B) {
	for _, tt := range filterBenchStringTests {
	    for i := 0; i < b.N; i++ {
			FilterString(tt.input, func(in string) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchIntTests = []struct {
	name     string
	input    []int
	equalTo  int
	expected []int
}{

	{
		name: "basic example",
		input: []int{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int{
            2,
        },
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
		expected: []int{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterInt(b *testing.B) {
	for _, tt := range filterBenchIntTests {
	    for i := 0; i < b.N; i++ {
			FilterInt(tt.input, func(in int) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchInt8Tests = []struct {
	name     string
	input    []int8
	equalTo  int8
	expected []int8
}{

	{
		name: "basic example",
		input: []int8{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int8{
            2,
        },
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
		expected: []int8{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterInt8(b *testing.B) {
	for _, tt := range filterBenchInt8Tests {
	    for i := 0; i < b.N; i++ {
			FilterInt8(tt.input, func(in int8) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchInt16Tests = []struct {
	name     string
	input    []int16
	equalTo  int16
	expected []int16
}{

	{
		name: "basic example",
		input: []int16{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int16{
            2,
        },
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
		expected: []int16{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterInt16(b *testing.B) {
	for _, tt := range filterBenchInt16Tests {
	    for i := 0; i < b.N; i++ {
			FilterInt16(tt.input, func(in int16) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchInt32Tests = []struct {
	name     string
	input    []int32
	equalTo  int32
	expected []int32
}{

	{
		name: "basic example",
		input: []int32{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int32{
            2,
        },
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
		expected: []int32{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterInt32(b *testing.B) {
	for _, tt := range filterBenchInt32Tests {
	    for i := 0; i < b.N; i++ {
			FilterInt32(tt.input, func(in int32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchInt64Tests = []struct {
	name     string
	input    []int64
	equalTo  int64
	expected []int64
}{

	{
		name: "basic example",
		input: []int64{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []int64{
            2,
        },
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
		expected: []int64{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterInt64(b *testing.B) {
	for _, tt := range filterBenchInt64Tests {
	    for i := 0; i < b.N; i++ {
			FilterInt64(tt.input, func(in int64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchUintTests = []struct {
	name     string
	input    []uint
	equalTo  uint
	expected []uint
}{

	{
		name: "basic example",
		input: []uint{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint{
            2,
        },
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
		expected: []uint{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterUint(b *testing.B) {
	for _, tt := range filterBenchUintTests {
	    for i := 0; i < b.N; i++ {
			FilterUint(tt.input, func(in uint) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchUint8Tests = []struct {
	name     string
	input    []uint8
	equalTo  uint8
	expected []uint8
}{

	{
		name: "basic example",
		input: []uint8{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint8{
            2,
        },
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
		expected: []uint8{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterUint8(b *testing.B) {
	for _, tt := range filterBenchUint8Tests {
	    for i := 0; i < b.N; i++ {
			FilterUint8(tt.input, func(in uint8) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchUint16Tests = []struct {
	name     string
	input    []uint16
	equalTo  uint16
	expected []uint16
}{

	{
		name: "basic example",
		input: []uint16{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint16{
            2,
        },
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
		expected: []uint16{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterUint16(b *testing.B) {
	for _, tt := range filterBenchUint16Tests {
	    for i := 0; i < b.N; i++ {
			FilterUint16(tt.input, func(in uint16) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchUint32Tests = []struct {
	name     string
	input    []uint32
	equalTo  uint32
	expected []uint32
}{

	{
		name: "basic example",
		input: []uint32{
			1,
			2,
			3,
		},
        equalTo: 2,
		expected: []uint32{
            2,
        },
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
		expected: []uint32{
            2,
            2,
        },
	},
}

func BenchmarkTestFilterUint32(b *testing.B) {
	for _, tt := range filterBenchUint32Tests {
	    for i := 0; i < b.N; i++ {
			FilterUint32(tt.input, func(in uint32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchFloat32Tests = []struct {
	name     string
	input    []float32
	equalTo  float32
	expected []float32
}{

	{
		name: "basic example",
		input: []float32{
			1.1,
			2.2,
			3.3,
		},
        equalTo: 2.2,
		expected: []float32{
            2.2,
        },
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
		expected: []float32{
            2.2,
            2.2,
        },
	},
}

func BenchmarkTestFilterFloat32(b *testing.B) {
	for _, tt := range filterBenchFloat32Tests {
	    for i := 0; i < b.N; i++ {
			FilterFloat32(tt.input, func(in float32) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchFloat64Tests = []struct {
	name     string
	input    []float64
	equalTo  float64
	expected []float64
}{

	{
		name: "basic example",
		input: []float64{
			1.1,
			2.2,
			3.3,
		},
        equalTo: 2.2,
		expected: []float64{
            2.2,
        },
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
		expected: []float64{
            2.2,
            2.2,
        },
	},
}

func BenchmarkTestFilterFloat64(b *testing.B) {
	for _, tt := range filterBenchFloat64Tests {
	    for i := 0; i < b.N; i++ {
			FilterFloat64(tt.input, func(in float64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchComplex64Tests = []struct {
	name     string
	input    []complex64
	equalTo  complex64
	expected []complex64
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.2+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: []complex64{
            (4.8+3.14i),
        },
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
		expected: []complex64{
            (4.8+3.14i),
            (4.8+3.14i),
        },
	},
}

func BenchmarkTestFilterComplex64(b *testing.B) {
	for _, tt := range filterBenchComplex64Tests {
	    for i := 0; i < b.N; i++ {
			FilterComplex64(tt.input, func(in complex64) bool {
				return in == tt.equalTo
			})
		}
	}

}

var filterBenchComplex128Tests = []struct {
	name     string
	input    []complex128
	equalTo  complex128
	expected []complex128
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
			(7.199999999999999+3.14i),
		},
        equalTo: (4.8+3.14i),
		expected: []complex128{
            (4.8+3.14i),
        },
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
		expected: []complex128{
            (4.8+3.14i),
            (4.8+3.14i),
        },
	},
}

func BenchmarkTestFilterComplex128(b *testing.B) {
	for _, tt := range filterBenchComplex128Tests {
	    for i := 0; i < b.N; i++ {
			FilterComplex128(tt.input, func(in complex128) bool {
				return in == tt.equalTo
			})
		}
	}

}
