package slice

import (
	"testing"
)
		
var uniqueBenchStringTests = []struct {
	name     string
	input    []string
	expected []string
}{
	{
		name: "no duplicates",
		input: []string{
			"test1",
			"test2",
		},
		expected: []string{
			"test1",
			"test2",
		},
	},
	{
		name: "one duplicates",
		input: []string{
			"test1",
			"test1",
			"test2",
		},
		expected: []string{
			"test1",
			"test2",
		},
	},
	{
		name: "multiple duplicates",
		input: []string{
			"test1",
			"test1",
			"test2",
			"test2",
			"test2",
			"test1",
			"test1",
			"test2",
			"test1",
		},
		expected: []string{
			"test1",
			"test2",
		},
	},
}

func BenchmarkTestUniqueString(b *testing.B) {
	for _, tt := range uniqueBenchStringTests {
		for i := 0; i < b.N; i++ {
			UniqueString(tt.input)
		}
	}
}

var uniqueBenchIntTests = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name: "no duplicates",
		input: []int{
			1,
			2,
		},
		expected: []int{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int{
			1,
			1,
			2,
		},
		expected: []int{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueInt(b *testing.B) {
	for _, tt := range uniqueBenchIntTests {
		for i := 0; i < b.N; i++ {
			UniqueInt(tt.input)
		}
	}
}

var uniqueBenchInt8Tests = []struct {
	name     string
	input    []int8
	expected []int8
}{
	{
		name: "no duplicates",
		input: []int8{
			1,
			2,
		},
		expected: []int8{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int8{
			1,
			1,
			2,
		},
		expected: []int8{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int8{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int8{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueInt8(b *testing.B) {
	for _, tt := range uniqueBenchInt8Tests {
		for i := 0; i < b.N; i++ {
			UniqueInt8(tt.input)
		}
	}
}

var uniqueBenchInt16Tests = []struct {
	name     string
	input    []int16
	expected []int16
}{
	{
		name: "no duplicates",
		input: []int16{
			1,
			2,
		},
		expected: []int16{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int16{
			1,
			1,
			2,
		},
		expected: []int16{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int16{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int16{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueInt16(b *testing.B) {
	for _, tt := range uniqueBenchInt16Tests {
		for i := 0; i < b.N; i++ {
			UniqueInt16(tt.input)
		}
	}
}

var uniqueBenchInt32Tests = []struct {
	name     string
	input    []int32
	expected []int32
}{
	{
		name: "no duplicates",
		input: []int32{
			1,
			2,
		},
		expected: []int32{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int32{
			1,
			1,
			2,
		},
		expected: []int32{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int32{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int32{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueInt32(b *testing.B) {
	for _, tt := range uniqueBenchInt32Tests {
		for i := 0; i < b.N; i++ {
			UniqueInt32(tt.input)
		}
	}
}

var uniqueBenchInt64Tests = []struct {
	name     string
	input    []int64
	expected []int64
}{
	{
		name: "no duplicates",
		input: []int64{
			1,
			2,
		},
		expected: []int64{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []int64{
			1,
			1,
			2,
		},
		expected: []int64{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []int64{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []int64{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueInt64(b *testing.B) {
	for _, tt := range uniqueBenchInt64Tests {
		for i := 0; i < b.N; i++ {
			UniqueInt64(tt.input)
		}
	}
}

var uniqueBenchUintTests = []struct {
	name     string
	input    []uint
	expected []uint
}{
	{
		name: "no duplicates",
		input: []uint{
			1,
			2,
		},
		expected: []uint{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint{
			1,
			1,
			2,
		},
		expected: []uint{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueUint(b *testing.B) {
	for _, tt := range uniqueBenchUintTests {
		for i := 0; i < b.N; i++ {
			UniqueUint(tt.input)
		}
	}
}

var uniqueBenchUint8Tests = []struct {
	name     string
	input    []uint8
	expected []uint8
}{
	{
		name: "no duplicates",
		input: []uint8{
			1,
			2,
		},
		expected: []uint8{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint8{
			1,
			1,
			2,
		},
		expected: []uint8{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint8{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint8{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueUint8(b *testing.B) {
	for _, tt := range uniqueBenchUint8Tests {
		for i := 0; i < b.N; i++ {
			UniqueUint8(tt.input)
		}
	}
}

var uniqueBenchUint16Tests = []struct {
	name     string
	input    []uint16
	expected []uint16
}{
	{
		name: "no duplicates",
		input: []uint16{
			1,
			2,
		},
		expected: []uint16{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint16{
			1,
			1,
			2,
		},
		expected: []uint16{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint16{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint16{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueUint16(b *testing.B) {
	for _, tt := range uniqueBenchUint16Tests {
		for i := 0; i < b.N; i++ {
			UniqueUint16(tt.input)
		}
	}
}

var uniqueBenchUint32Tests = []struct {
	name     string
	input    []uint32
	expected []uint32
}{
	{
		name: "no duplicates",
		input: []uint32{
			1,
			2,
		},
		expected: []uint32{
			1,
			2,
		},
	},
	{
		name: "one duplicates",
		input: []uint32{
			1,
			1,
			2,
		},
		expected: []uint32{
			1,
			2,
		},
	},
	{
		name: "multiple duplicates",
		input: []uint32{
			1,
			1,
			2,
			2,
			2,
			1,
			1,
			2,
			1,
		},
		expected: []uint32{
			1,
			2,
		},
	},
}

func BenchmarkTestUniqueUint32(b *testing.B) {
	for _, tt := range uniqueBenchUint32Tests {
		for i := 0; i < b.N; i++ {
			UniqueUint32(tt.input)
		}
	}
}

var uniqueBenchFloat32Tests = []struct {
	name     string
	input    []float32
	expected []float32
}{
	{
		name: "no duplicates",
		input: []float32{
			1.1,
			2.2,
		},
		expected: []float32{
			1.1,
			2.2,
		},
	},
	{
		name: "one duplicates",
		input: []float32{
			1.1,
			1.1,
			2.2,
		},
		expected: []float32{
			1.1,
			2.2,
		},
	},
	{
		name: "multiple duplicates",
		input: []float32{
			1.1,
			1.1,
			2.2,
			2.2,
			2.2,
			1.1,
			1.1,
			2.2,
			1.1,
		},
		expected: []float32{
			1.1,
			2.2,
		},
	},
}

func BenchmarkTestUniqueFloat32(b *testing.B) {
	for _, tt := range uniqueBenchFloat32Tests {
		for i := 0; i < b.N; i++ {
			UniqueFloat32(tt.input)
		}
	}
}

var uniqueBenchFloat64Tests = []struct {
	name     string
	input    []float64
	expected []float64
}{
	{
		name: "no duplicates",
		input: []float64{
			1.1,
			2.2,
		},
		expected: []float64{
			1.1,
			2.2,
		},
	},
	{
		name: "one duplicates",
		input: []float64{
			1.1,
			1.1,
			2.2,
		},
		expected: []float64{
			1.1,
			2.2,
		},
	},
	{
		name: "multiple duplicates",
		input: []float64{
			1.1,
			1.1,
			2.2,
			2.2,
			2.2,
			1.1,
			1.1,
			2.2,
			1.1,
		},
		expected: []float64{
			1.1,
			2.2,
		},
	},
}

func BenchmarkTestUniqueFloat64(b *testing.B) {
	for _, tt := range uniqueBenchFloat64Tests {
		for i := 0; i < b.N; i++ {
			UniqueFloat64(tt.input)
		}
	}
}

var uniqueBenchComplex64Tests = []struct {
	name     string
	input    []complex64
	expected []complex64
}{
	{
		name: "no duplicates",
		input: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "one duplicates",
		input: []complex64{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "multiple duplicates",
		input: []complex64{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
		},
		expected: []complex64{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
}

func BenchmarkTestUniqueComplex64(b *testing.B) {
	for _, tt := range uniqueBenchComplex64Tests {
		for i := 0; i < b.N; i++ {
			UniqueComplex64(tt.input)
		}
	}
}

var uniqueBenchComplex128Tests = []struct {
	name     string
	input    []complex128
	expected []complex128
}{
	{
		name: "no duplicates",
		input: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "one duplicates",
		input: []complex128{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
		},
		expected: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
	{
		name: "multiple duplicates",
		input: []complex128{
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
			(2.4+3.14i),
			(4.8+3.14i),
			(2.4+3.14i),
		},
		expected: []complex128{
			(2.4+3.14i),
			(4.8+3.14i),
		},
	},
}

func BenchmarkTestUniqueComplex128(b *testing.B) {
	for _, tt := range uniqueBenchComplex128Tests {
		for i := 0; i < b.N; i++ {
			UniqueComplex128(tt.input)
		}
	}
}
