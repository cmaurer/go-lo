package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

		

// String Example
func ExampleRemoveString() {

    input := []string{
        "test1",
        "test2",
        "test3",
        "test2",
    }

    equalTo := string("test2")

    output := slice.RemoveString(input, func(in string) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Int Example
func ExampleRemoveInt() {

    input := []int{
        1,
        2,
        3,
        2,
    }

    equalTo := int(2)

    output := slice.RemoveInt(input, func(in int) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Int8 Example
func ExampleRemoveInt8() {

    input := []int8{
        1,
        2,
        3,
        2,
    }

    equalTo := int8(2)

    output := slice.RemoveInt8(input, func(in int8) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Int16 Example
func ExampleRemoveInt16() {

    input := []int16{
        1,
        2,
        3,
        2,
    }

    equalTo := int16(2)

    output := slice.RemoveInt16(input, func(in int16) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Int32 Example
func ExampleRemoveInt32() {

    input := []int32{
        1,
        2,
        3,
        2,
    }

    equalTo := int32(2)

    output := slice.RemoveInt32(input, func(in int32) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Int64 Example
func ExampleRemoveInt64() {

    input := []int64{
        1,
        2,
        3,
        2,
    }

    equalTo := int64(2)

    output := slice.RemoveInt64(input, func(in int64) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Uint Example
func ExampleRemoveUint() {

    input := []uint{
        1,
        2,
        3,
        2,
    }

    equalTo := uint(2)

    output := slice.RemoveUint(input, func(in uint) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Uint8 Example
func ExampleRemoveUint8() {

    input := []uint8{
        1,
        2,
        3,
        2,
    }

    equalTo := uint8(2)

    output := slice.RemoveUint8(input, func(in uint8) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Uint16 Example
func ExampleRemoveUint16() {

    input := []uint16{
        1,
        2,
        3,
        2,
    }

    equalTo := uint16(2)

    output := slice.RemoveUint16(input, func(in uint16) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Uint32 Example
func ExampleRemoveUint32() {

    input := []uint32{
        1,
        2,
        3,
        2,
    }

    equalTo := uint32(2)

    output := slice.RemoveUint32(input, func(in uint32) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Float32 Example
func ExampleRemoveFloat32() {

    input := []float32{
        1.1,
        2.2,
        3.3,
        2.2,
    }

    equalTo := float32(2.2)

    output := slice.RemoveFloat32(input, func(in float32) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Float64 Example
func ExampleRemoveFloat64() {

    input := []float64{
        1.1,
        2.2,
        3.3,
        2.2,
    }

    equalTo := float64(2.2)

    output := slice.RemoveFloat64(input, func(in float64) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Complex64 Example
func ExampleRemoveComplex64() {

    input := []complex64{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.2+3.14i),
        (4.8+3.14i),
    }

    equalTo := complex64((4.8+3.14i))

    output := slice.RemoveComplex64(input, func(in complex64) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}

// Complex128 Example
func ExampleRemoveComplex128() {

    input := []complex128{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.199999999999999+3.14i),
        (4.8+3.14i),
    }

    equalTo := complex128((4.8+3.14i))

    output := slice.RemoveComplex128(input, func(in complex128) bool {
        return in == equalTo
    })

    fmt.Printf("%+v", output)

}