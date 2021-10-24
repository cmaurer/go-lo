package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

		
// string Example
func ExampleReduceString() {

	input := []string{
		"test1",
		"test2",
		"test3",
	}

	output := slice.ReduceString(input, func(acc string, elem string) string {
		return acc + elem
	}, "test0")
	fmt.Printf("%+v", output)
}

// int Example
func ExampleReduceInt() {

	input := []int{
		1,
		2,
		3,
	}

	output := slice.ReduceInt(input, func(acc int, elem int) int {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// int8 Example
func ExampleReduceInt8() {

	input := []int8{
		1,
		2,
		3,
	}

	output := slice.ReduceInt8(input, func(acc int8, elem int8) int8 {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// int16 Example
func ExampleReduceInt16() {

	input := []int16{
		1,
		2,
		3,
	}

	output := slice.ReduceInt16(input, func(acc int16, elem int16) int16 {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// int32 Example
func ExampleReduceInt32() {

	input := []int32{
		1,
		2,
		3,
	}

	output := slice.ReduceInt32(input, func(acc int32, elem int32) int32 {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// int64 Example
func ExampleReduceInt64() {

	input := []int64{
		1,
		2,
		3,
	}

	output := slice.ReduceInt64(input, func(acc int64, elem int64) int64 {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// uint Example
func ExampleReduceUint() {

	input := []uint{
		1,
		2,
		3,
	}

	output := slice.ReduceUint(input, func(acc uint, elem uint) uint {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// uint8 Example
func ExampleReduceUint8() {

	input := []uint8{
		1,
		2,
		3,
	}

	output := slice.ReduceUint8(input, func(acc uint8, elem uint8) uint8 {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// uint16 Example
func ExampleReduceUint16() {

	input := []uint16{
		1,
		2,
		3,
	}

	output := slice.ReduceUint16(input, func(acc uint16, elem uint16) uint16 {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// uint32 Example
func ExampleReduceUint32() {

	input := []uint32{
		1,
		2,
		3,
	}

	output := slice.ReduceUint32(input, func(acc uint32, elem uint32) uint32 {
		return acc + elem
	}, 0)
	fmt.Printf("%+v", output)
}

// float32 Example
func ExampleReduceFloat32() {

	input := []float32{
		1.1,
		2.2,
		3.3,
	}

	output := slice.ReduceFloat32(input, func(acc float32, elem float32) float32 {
		return acc + elem
	}, 0.0)
	fmt.Printf("%+v", output)
}

// float64 Example
func ExampleReduceFloat64() {

	input := []float64{
		1.1,
		2.2,
		3.3,
	}

	output := slice.ReduceFloat64(input, func(acc float64, elem float64) float64 {
		return acc + elem
	}, 0.0)
	fmt.Printf("%+v", output)
}

// complex64 Example
func ExampleReduceComplex64() {

	input := []complex64{
		(2.4+3.14i),
		(4.8+3.14i),
		(7.2+3.14i),
	}

	output := slice.ReduceComplex64(input, func(acc complex64, elem complex64) complex64 {
		return acc + elem
	}, (0+3.14i))
	fmt.Printf("%+v", output)
}

// complex128 Example
func ExampleReduceComplex128() {

	input := []complex128{
		(2.4+3.14i),
		(4.8+3.14i),
		(7.199999999999999+3.14i),
	}

	output := slice.ReduceComplex128(input, func(acc complex128, elem complex128) complex128 {
		return acc + elem
	}, (0+3.14i))
	fmt.Printf("%+v", output)
}
