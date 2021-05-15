package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
		


var mapStringTests = []struct {
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

func TestMapString(t *testing.T) {
	for _, tt := range mapStringTests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapString(tt.input, func(in string) string {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapIntTests = []struct {
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

func TestMapInt(t *testing.T) {
	for _, tt := range mapIntTests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapInt(tt.input, func(in int) int {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapInt8Tests = []struct {
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

func TestMapInt8(t *testing.T) {
	for _, tt := range mapInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapInt8(tt.input, func(in int8) int8 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapInt16Tests = []struct {
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

func TestMapInt16(t *testing.T) {
	for _, tt := range mapInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapInt16(tt.input, func(in int16) int16 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapInt32Tests = []struct {
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

func TestMapInt32(t *testing.T) {
	for _, tt := range mapInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapInt32(tt.input, func(in int32) int32 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapInt64Tests = []struct {
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

func TestMapInt64(t *testing.T) {
	for _, tt := range mapInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapInt64(tt.input, func(in int64) int64 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapUintTests = []struct {
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

func TestMapUint(t *testing.T) {
	for _, tt := range mapUintTests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapUint(tt.input, func(in uint) uint {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapUint8Tests = []struct {
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

func TestMapUint8(t *testing.T) {
	for _, tt := range mapUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapUint8(tt.input, func(in uint8) uint8 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapUint16Tests = []struct {
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

func TestMapUint16(t *testing.T) {
	for _, tt := range mapUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapUint16(tt.input, func(in uint16) uint16 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapUint32Tests = []struct {
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

func TestMapUint32(t *testing.T) {
	for _, tt := range mapUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapUint32(tt.input, func(in uint32) uint32 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapFloat32Tests = []struct {
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

func TestMapFloat32(t *testing.T) {
	for _, tt := range mapFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapFloat32(tt.input, func(in float32) float32 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapFloat64Tests = []struct {
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

func TestMapFloat64(t *testing.T) {
	for _, tt := range mapFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapFloat64(tt.input, func(in float64) float64 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapComplex64Tests = []struct {
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

func TestMapComplex64(t *testing.T) {
	for _, tt := range mapComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapComplex64(tt.input, func(in complex64) complex64 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}



var mapComplex128Tests = []struct {
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

func TestMapComplex128(t *testing.T) {
	for _, tt := range mapComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := MapComplex128(tt.input, func(in complex128) complex128 {
                return in + in
            })
			assert.Equal(t, output, tt.expected)
		})
	}

}
