package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var flattenStringTests = []struct {
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

func TestFlattenString(t *testing.T) {
	for _, tt := range flattenStringTests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenString(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenIntTests = []struct {
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

func TestFlattenInt(t *testing.T) {
	for _, tt := range flattenIntTests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenInt(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenInt8Tests = []struct {
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

func TestFlattenInt8(t *testing.T) {
	for _, tt := range flattenInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenInt8(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenInt16Tests = []struct {
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

func TestFlattenInt16(t *testing.T) {
	for _, tt := range flattenInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenInt16(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenInt32Tests = []struct {
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

func TestFlattenInt32(t *testing.T) {
	for _, tt := range flattenInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenInt32(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenInt64Tests = []struct {
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

func TestFlattenInt64(t *testing.T) {
	for _, tt := range flattenInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenInt64(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenUintTests = []struct {
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

func TestFlattenUint(t *testing.T) {
	for _, tt := range flattenUintTests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenUint(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenUint8Tests = []struct {
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

func TestFlattenUint8(t *testing.T) {
	for _, tt := range flattenUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenUint8(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenUint16Tests = []struct {
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

func TestFlattenUint16(t *testing.T) {
	for _, tt := range flattenUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenUint16(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenUint32Tests = []struct {
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

func TestFlattenUint32(t *testing.T) {
	for _, tt := range flattenUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenUint32(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenFloat32Tests = []struct {
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

func TestFlattenFloat32(t *testing.T) {
	for _, tt := range flattenFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenFloat32(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenFloat64Tests = []struct {
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

func TestFlattenFloat64(t *testing.T) {
	for _, tt := range flattenFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenFloat64(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenComplex64Tests = []struct {
	name     string
	input1   []complex64
	input2   []complex64
	expected []complex64
}{

	{
		name: "basic example",
		input1: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
		},
		input2: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
		},
		expected: []complex64{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.2 + 3.14i),
		},
	},
}

func TestFlattenComplex64(t *testing.T) {
	for _, tt := range flattenComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenComplex64(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}

var flattenComplex128Tests = []struct {
	name     string
	input1   []complex128
	input2   []complex128
	expected []complex128
}{

	{
		name: "basic example",
		input1: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
		},
		input2: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
		},
		expected: []complex128{
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
			(2.4 + 3.14i),
			(4.8 + 3.14i),
			(7.199999999999999 + 3.14i),
		},
	},
}

func TestFlattenComplex128(t *testing.T) {
	for _, tt := range flattenComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FlattenComplex128(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, output)
		})
	}

}
