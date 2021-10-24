package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

// String Example
func ExampleMapString() {

	input := []string{
		"test1",
		"test2",
		"test3",
		"test2",
	}

	output := slice.MapString(input, func(in string) string {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Int Example
func ExampleMapInt() {

	input := []int{
		1,
		2,
		3,
		2,
	}

	output := slice.MapInt(input, func(in int) int {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Int8 Example
func ExampleMapInt8() {

	input := []int8{
		1,
		2,
		3,
		2,
	}

	output := slice.MapInt8(input, func(in int8) int8 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Int16 Example
func ExampleMapInt16() {

	input := []int16{
		1,
		2,
		3,
		2,
	}

	output := slice.MapInt16(input, func(in int16) int16 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Int32 Example
func ExampleMapInt32() {

	input := []int32{
		1,
		2,
		3,
		2,
	}

	output := slice.MapInt32(input, func(in int32) int32 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Int64 Example
func ExampleMapInt64() {

	input := []int64{
		1,
		2,
		3,
		2,
	}

	output := slice.MapInt64(input, func(in int64) int64 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Uint Example
func ExampleMapUint() {

	input := []uint{
		1,
		2,
		3,
		2,
	}

	output := slice.MapUint(input, func(in uint) uint {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Uint8 Example
func ExampleMapUint8() {

	input := []uint8{
		1,
		2,
		3,
		2,
	}

	output := slice.MapUint8(input, func(in uint8) uint8 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Uint16 Example
func ExampleMapUint16() {

	input := []uint16{
		1,
		2,
		3,
		2,
	}

	output := slice.MapUint16(input, func(in uint16) uint16 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Uint32 Example
func ExampleMapUint32() {

	input := []uint32{
		1,
		2,
		3,
		2,
	}

	output := slice.MapUint32(input, func(in uint32) uint32 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Float32 Example
func ExampleMapFloat32() {

	input := []float32{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	output := slice.MapFloat32(input, func(in float32) float32 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Float64 Example
func ExampleMapFloat64() {

	input := []float64{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	output := slice.MapFloat64(input, func(in float64) float64 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Complex64 Example
func ExampleMapComplex64() {

	input := []complex64{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.2 + 3.14i),
		(4.8 + 3.14i),
	}

	output := slice.MapComplex64(input, func(in complex64) complex64 {
		return in + in
	})

	fmt.Printf("%+v", output)

}

// Complex128 Example
func ExampleMapComplex128() {

	input := []complex128{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.199999999999999 + 3.14i),
		(4.8 + 3.14i),
	}

	output := slice.MapComplex128(input, func(in complex128) complex128 {
		return in + in
	})

	fmt.Printf("%+v", output)

}
