package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

// String Example
func ExampleEachString() {

	input := []string{
		"test1",
		"test2",
		"test3",
		"test2",
	}

	expectedCount := 0

	slice.EachString(input, func(in string) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Int Example
func ExampleEachInt() {

	input := []int{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachInt(input, func(in int) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Int8 Example
func ExampleEachInt8() {

	input := []int8{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachInt8(input, func(in int8) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Int16 Example
func ExampleEachInt16() {

	input := []int16{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachInt16(input, func(in int16) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Int32 Example
func ExampleEachInt32() {

	input := []int32{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachInt32(input, func(in int32) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Int64 Example
func ExampleEachInt64() {

	input := []int64{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachInt64(input, func(in int64) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Uint Example
func ExampleEachUint() {

	input := []uint{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachUint(input, func(in uint) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Uint8 Example
func ExampleEachUint8() {

	input := []uint8{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachUint8(input, func(in uint8) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Uint16 Example
func ExampleEachUint16() {

	input := []uint16{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachUint16(input, func(in uint16) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Uint32 Example
func ExampleEachUint32() {

	input := []uint32{
		1,
		2,
		3,
		2,
	}

	expectedCount := 0

	slice.EachUint32(input, func(in uint32) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Float32 Example
func ExampleEachFloat32() {

	input := []float32{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	expectedCount := 0

	slice.EachFloat32(input, func(in float32) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Float64 Example
func ExampleEachFloat64() {

	input := []float64{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	expectedCount := 0

	slice.EachFloat64(input, func(in float64) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Complex64 Example
func ExampleEachComplex64() {

	input := []complex64{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.2 + 3.14i),
		(4.8 + 3.14i),
	}

	expectedCount := 0

	slice.EachComplex64(input, func(in complex64) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}

// Complex128 Example
func ExampleEachComplex128() {

	input := []complex128{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.199999999999999 + 3.14i),
		(4.8 + 3.14i),
	}

	expectedCount := 0

	slice.EachComplex128(input, func(in complex128) {
		expectedCount = expectedCount + 1
	})

	fmt.Printf("%+v", expectedCount)

}
