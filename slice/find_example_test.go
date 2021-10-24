package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

// String Example
func ExampleFindString() {

	input := []string{
		"test1",
		"test2",
		"test3",
		"test2",
	}

	equalTo := string("test8")

	output := slice.FindString(input, func(in string) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Int Example
func ExampleFindInt() {

	input := []int{
		1,
		2,
		3,
		2,
	}

	equalTo := int(8)

	output := slice.FindInt(input, func(in int) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Int8 Example
func ExampleFindInt8() {

	input := []int8{
		1,
		2,
		3,
		2,
	}

	equalTo := int8(8)

	output := slice.FindInt8(input, func(in int8) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Int16 Example
func ExampleFindInt16() {

	input := []int16{
		1,
		2,
		3,
		2,
	}

	equalTo := int16(8)

	output := slice.FindInt16(input, func(in int16) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Int32 Example
func ExampleFindInt32() {

	input := []int32{
		1,
		2,
		3,
		2,
	}

	equalTo := int32(8)

	output := slice.FindInt32(input, func(in int32) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Int64 Example
func ExampleFindInt64() {

	input := []int64{
		1,
		2,
		3,
		2,
	}

	equalTo := int64(8)

	output := slice.FindInt64(input, func(in int64) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Uint Example
func ExampleFindUint() {

	input := []uint{
		1,
		2,
		3,
		2,
	}

	equalTo := uint(8)

	output := slice.FindUint(input, func(in uint) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Uint8 Example
func ExampleFindUint8() {

	input := []uint8{
		1,
		2,
		3,
		2,
	}

	equalTo := uint8(8)

	output := slice.FindUint8(input, func(in uint8) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Uint16 Example
func ExampleFindUint16() {

	input := []uint16{
		1,
		2,
		3,
		2,
	}

	equalTo := uint16(8)

	output := slice.FindUint16(input, func(in uint16) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Uint32 Example
func ExampleFindUint32() {

	input := []uint32{
		1,
		2,
		3,
		2,
	}

	equalTo := uint32(8)

	output := slice.FindUint32(input, func(in uint32) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Float32 Example
func ExampleFindFloat32() {

	input := []float32{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	equalTo := float32(8.8)

	output := slice.FindFloat32(input, func(in float32) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Float64 Example
func ExampleFindFloat64() {

	input := []float64{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	equalTo := float64(8.8)

	output := slice.FindFloat64(input, func(in float64) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Complex64 Example
func ExampleFindComplex64() {

	input := []complex64{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.2 + 3.14i),
		(4.8 + 3.14i),
	}

	equalTo := complex64((19.2 + 3.14i))

	output := slice.FindComplex64(input, func(in complex64) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}

// Complex128 Example
func ExampleFindComplex128() {

	input := []complex128{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.199999999999999 + 3.14i),
		(4.8 + 3.14i),
	}

	equalTo := complex128((19.2 + 3.14i))

	output := slice.FindComplex128(input, func(in complex128) bool {
		return in == equalTo
	})

	fmt.Printf("%+v", output)

}
