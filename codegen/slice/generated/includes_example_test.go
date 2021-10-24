package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

		
// String Example
func ExampleIncludesString() {

    input := []string{
        "test1",
        "test2",
        "test3",
        "test2",
    }

	included := string("test3")

	output := slice.IncludesString(input, included)

    fmt.Printf("%+v", output)

}
// Int Example
func ExampleIncludesInt() {

    input := []int{
        1,
        2,
        3,
        2,
    }

	included := int(3)

	output := slice.IncludesInt(input, included)

    fmt.Printf("%+v", output)

}
// Int8 Example
func ExampleIncludesInt8() {

    input := []int8{
        1,
        2,
        3,
        2,
    }

	included := int8(3)

	output := slice.IncludesInt8(input, included)

    fmt.Printf("%+v", output)

}
// Int16 Example
func ExampleIncludesInt16() {

    input := []int16{
        1,
        2,
        3,
        2,
    }

	included := int16(3)

	output := slice.IncludesInt16(input, included)

    fmt.Printf("%+v", output)

}
// Int32 Example
func ExampleIncludesInt32() {

    input := []int32{
        1,
        2,
        3,
        2,
    }

	included := int32(3)

	output := slice.IncludesInt32(input, included)

    fmt.Printf("%+v", output)

}
// Int64 Example
func ExampleIncludesInt64() {

    input := []int64{
        1,
        2,
        3,
        2,
    }

	included := int64(3)

	output := slice.IncludesInt64(input, included)

    fmt.Printf("%+v", output)

}
// Uint Example
func ExampleIncludesUint() {

    input := []uint{
        1,
        2,
        3,
        2,
    }

	included := uint(3)

	output := slice.IncludesUint(input, included)

    fmt.Printf("%+v", output)

}
// Uint8 Example
func ExampleIncludesUint8() {

    input := []uint8{
        1,
        2,
        3,
        2,
    }

	included := uint8(3)

	output := slice.IncludesUint8(input, included)

    fmt.Printf("%+v", output)

}
// Uint16 Example
func ExampleIncludesUint16() {

    input := []uint16{
        1,
        2,
        3,
        2,
    }

	included := uint16(3)

	output := slice.IncludesUint16(input, included)

    fmt.Printf("%+v", output)

}
// Uint32 Example
func ExampleIncludesUint32() {

    input := []uint32{
        1,
        2,
        3,
        2,
    }

	included := uint32(3)

	output := slice.IncludesUint32(input, included)

    fmt.Printf("%+v", output)

}
// Float32 Example
func ExampleIncludesFloat32() {

    input := []float32{
        1.1,
        2.2,
        3.3,
        2.2,
    }

	included := float32(3.3)

	output := slice.IncludesFloat32(input, included)

    fmt.Printf("%+v", output)

}
// Float64 Example
func ExampleIncludesFloat64() {

    input := []float64{
        1.1,
        2.2,
        3.3,
        2.2,
    }

	included := float64(3.3)

	output := slice.IncludesFloat64(input, included)

    fmt.Printf("%+v", output)

}
// Complex64 Example
func ExampleIncludesComplex64() {

    input := []complex64{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.2+3.14i),
        (4.8+3.14i),
    }

	included := complex64((7.2+3.14i))

	output := slice.IncludesComplex64(input, included)

    fmt.Printf("%+v", output)

}
// Complex128 Example
func ExampleIncludesComplex128() {

    input := []complex128{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.199999999999999+3.14i),
        (4.8+3.14i),
    }

	included := complex128((7.199999999999999+3.14i))

	output := slice.IncludesComplex128(input, included)

    fmt.Printf("%+v", output)

}