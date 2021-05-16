package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var eachStringTests = []struct {
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

func TestEachString(t *testing.T) {
	for _, tt := range eachStringTests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachString(tt.input, func(in string) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachIntTests = []struct {
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

func TestEachInt(t *testing.T) {
	for _, tt := range eachIntTests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachInt(tt.input, func(in int) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachInt8Tests = []struct {
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

func TestEachInt8(t *testing.T) {
	for _, tt := range eachInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachInt8(tt.input, func(in int8) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachInt16Tests = []struct {
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

func TestEachInt16(t *testing.T) {
	for _, tt := range eachInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachInt16(tt.input, func(in int16) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachInt32Tests = []struct {
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

func TestEachInt32(t *testing.T) {
	for _, tt := range eachInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachInt32(tt.input, func(in int32) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachInt64Tests = []struct {
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

func TestEachInt64(t *testing.T) {
	for _, tt := range eachInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachInt64(tt.input, func(in int64) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachUintTests = []struct {
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

func TestEachUint(t *testing.T) {
	for _, tt := range eachUintTests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachUint(tt.input, func(in uint) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachUint8Tests = []struct {
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

func TestEachUint8(t *testing.T) {
	for _, tt := range eachUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachUint8(tt.input, func(in uint8) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachUint16Tests = []struct {
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

func TestEachUint16(t *testing.T) {
	for _, tt := range eachUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachUint16(tt.input, func(in uint16) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachUint32Tests = []struct {
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

func TestEachUint32(t *testing.T) {
	for _, tt := range eachUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachUint32(tt.input, func(in uint32) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachFloat32Tests = []struct {
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

func TestEachFloat32(t *testing.T) {
	for _, tt := range eachFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachFloat32(tt.input, func(in float32) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachFloat64Tests = []struct {
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

func TestEachFloat64(t *testing.T) {
	for _, tt := range eachFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachFloat64(tt.input, func(in float64) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachComplex64Tests = []struct {
	name     string
	input    []complex64
	expected int
}{

	{
		name: "basic example",
		input: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
		},
		expected: 3,
	},
}

func TestEachComplex64(t *testing.T) {
	for _, tt := range eachComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachComplex64(tt.input, func(in complex64) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}

var eachComplex128Tests = []struct {
	name     string
	input    []complex128
	expected int
}{

	{
		name: "basic example",
		input: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
		},
		expected: 3,
	},
}

func TestEachComplex128(t *testing.T) {
	for _, tt := range eachComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCount := 0
			EachComplex128(tt.input, func(in complex128) {
				expectedCount = expectedCount + 1
			})
			assert.Equal(t, expectedCount, tt.expected)
		})
	}

}
