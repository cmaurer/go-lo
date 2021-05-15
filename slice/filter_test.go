package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
		
var filterStringTests = []struct {
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

func TestFilterString(t *testing.T) {
	for _, tt := range filterStringTests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterString(tt.input, func(in string) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterIntTests = []struct {
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

func TestFilterInt(t *testing.T) {
	for _, tt := range filterIntTests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterInt(tt.input, func(in int) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterInt8Tests = []struct {
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

func TestFilterInt8(t *testing.T) {
	for _, tt := range filterInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterInt8(tt.input, func(in int8) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterInt16Tests = []struct {
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

func TestFilterInt16(t *testing.T) {
	for _, tt := range filterInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterInt16(tt.input, func(in int16) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterInt32Tests = []struct {
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

func TestFilterInt32(t *testing.T) {
	for _, tt := range filterInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterInt32(tt.input, func(in int32) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterInt64Tests = []struct {
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

func TestFilterInt64(t *testing.T) {
	for _, tt := range filterInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterInt64(tt.input, func(in int64) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterUintTests = []struct {
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

func TestFilterUint(t *testing.T) {
	for _, tt := range filterUintTests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterUint(tt.input, func(in uint) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterUint8Tests = []struct {
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

func TestFilterUint8(t *testing.T) {
	for _, tt := range filterUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterUint8(tt.input, func(in uint8) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterUint16Tests = []struct {
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

func TestFilterUint16(t *testing.T) {
	for _, tt := range filterUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterUint16(tt.input, func(in uint16) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterUint32Tests = []struct {
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

func TestFilterUint32(t *testing.T) {
	for _, tt := range filterUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterUint32(tt.input, func(in uint32) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterFloat32Tests = []struct {
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

func TestFilterFloat32(t *testing.T) {
	for _, tt := range filterFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterFloat32(tt.input, func(in float32) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterFloat64Tests = []struct {
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

func TestFilterFloat64(t *testing.T) {
	for _, tt := range filterFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterFloat64(tt.input, func(in float64) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterComplex64Tests = []struct {
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

func TestFilterComplex64(t *testing.T) {
	for _, tt := range filterComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterComplex64(tt.input, func(in complex64) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}

var filterComplex128Tests = []struct {
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

func TestFilterComplex128(t *testing.T) {
	for _, tt := range filterComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := FilterComplex128(tt.input, func(in complex128) bool {
                return in == tt.equalTo
            })
			assert.Equal(t, tt.expected, output)
		})
	}

}
