package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
		
var removeStringTests = []struct {
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

func TestRemoveString(t *testing.T) {
	for _, tt := range removeStringTests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveString(tt.input, func(in string) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeIntTests = []struct {
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

func TestRemoveInt(t *testing.T) {
	for _, tt := range removeIntTests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveInt(tt.input, func(in int) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeInt8Tests = []struct {
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

func TestRemoveInt8(t *testing.T) {
	for _, tt := range removeInt8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveInt8(tt.input, func(in int8) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeInt16Tests = []struct {
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

func TestRemoveInt16(t *testing.T) {
	for _, tt := range removeInt16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveInt16(tt.input, func(in int16) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeInt32Tests = []struct {
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

func TestRemoveInt32(t *testing.T) {
	for _, tt := range removeInt32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveInt32(tt.input, func(in int32) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeInt64Tests = []struct {
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

func TestRemoveInt64(t *testing.T) {
	for _, tt := range removeInt64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveInt64(tt.input, func(in int64) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeUintTests = []struct {
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

func TestRemoveUint(t *testing.T) {
	for _, tt := range removeUintTests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveUint(tt.input, func(in uint) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeUint8Tests = []struct {
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

func TestRemoveUint8(t *testing.T) {
	for _, tt := range removeUint8Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveUint8(tt.input, func(in uint8) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeUint16Tests = []struct {
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

func TestRemoveUint16(t *testing.T) {
	for _, tt := range removeUint16Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveUint16(tt.input, func(in uint16) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeUint32Tests = []struct {
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

func TestRemoveUint32(t *testing.T) {
	for _, tt := range removeUint32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveUint32(tt.input, func(in uint32) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeFloat32Tests = []struct {
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

func TestRemoveFloat32(t *testing.T) {
	for _, tt := range removeFloat32Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveFloat32(tt.input, func(in float32) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeFloat64Tests = []struct {
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

func TestRemoveFloat64(t *testing.T) {
	for _, tt := range removeFloat64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveFloat64(tt.input, func(in float64) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeComplex64Tests = []struct {
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

func TestRemoveComplex64(t *testing.T) {
	for _, tt := range removeComplex64Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveComplex64(tt.input, func(in complex64) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}

var removeComplex128Tests = []struct {
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

func TestRemoveComplex128(t *testing.T) {
	for _, tt := range removeComplex128Tests {
		t.Run(tt.name, func(t *testing.T) {
            output := RemoveComplex128(tt.input, func(in complex128) bool {
                return in == tt.equalTo
            })
            assert.Equal(t, tt.expected, output)
		})
	}

}
