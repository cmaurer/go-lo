package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

// String Example
func ExampleUniqueString() {

	input := []string{
		"test1",
		"test2",
		"test3",
		"test2",
	}

	output := slice.UniqueString(input)

	fmt.Printf("%+v", output)

}

// Int Example
func ExampleUniqueInt() {

	input := []int{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueInt(input)

	fmt.Printf("%+v", output)

}

// Int8 Example
func ExampleUniqueInt8() {

	input := []int8{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueInt8(input)

	fmt.Printf("%+v", output)

}

// Int16 Example
func ExampleUniqueInt16() {

	input := []int16{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueInt16(input)

	fmt.Printf("%+v", output)

}

// Int32 Example
func ExampleUniqueInt32() {

	input := []int32{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueInt32(input)

	fmt.Printf("%+v", output)

}

// Int64 Example
func ExampleUniqueInt64() {

	input := []int64{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueInt64(input)

	fmt.Printf("%+v", output)

}

// Uint Example
func ExampleUniqueUint() {

	input := []uint{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueUint(input)

	fmt.Printf("%+v", output)

}

// Uint8 Example
func ExampleUniqueUint8() {

	input := []uint8{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueUint8(input)

	fmt.Printf("%+v", output)

}

// Uint16 Example
func ExampleUniqueUint16() {

	input := []uint16{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueUint16(input)

	fmt.Printf("%+v", output)

}

// Uint32 Example
func ExampleUniqueUint32() {

	input := []uint32{
		1,
		2,
		3,
		2,
	}

	output := slice.UniqueUint32(input)

	fmt.Printf("%+v", output)

}

// Float32 Example
func ExampleUniqueFloat32() {

	input := []float32{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	output := slice.UniqueFloat32(input)

	fmt.Printf("%+v", output)

}

// Float64 Example
func ExampleUniqueFloat64() {

	input := []float64{
		1.1,
		2.2,
		3.3,
		2.2,
	}

	output := slice.UniqueFloat64(input)

	fmt.Printf("%+v", output)

}

// Complex64 Example
func ExampleUniqueComplex64() {

	input := []complex64{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.2 + 3.14i),
		(4.8 + 3.14i),
	}

	output := slice.UniqueComplex64(input)

	fmt.Printf("%+v", output)

}

// Complex128 Example
func ExampleUniqueComplex128() {

	input := []complex128{
		(2.4 + 3.14i),
		(4.8 + 3.14i),
		(7.199999999999999 + 3.14i),
		(4.8 + 3.14i),
	}

	output := slice.UniqueComplex128(input)

	fmt.Printf("%+v", output)

}
