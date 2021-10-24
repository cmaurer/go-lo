package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

		

// String Example
func ExampleIndexOfString() {

    input := []string{
        "test1",
        "test2",
        "test3",
        "test2",
    }
	element := string("test3")
	output := slice.IndexOfString(element, input)

    fmt.Printf("%+v", output)

}

// Int Example
func ExampleIndexOfInt() {

    input := []int{
        1,
        2,
        3,
        2,
    }
	element := int(3)
	output := slice.IndexOfInt(element, input)

    fmt.Printf("%+v", output)

}

// Int8 Example
func ExampleIndexOfInt8() {

    input := []int8{
        1,
        2,
        3,
        2,
    }
	element := int8(3)
	output := slice.IndexOfInt8(element, input)

    fmt.Printf("%+v", output)

}

// Int16 Example
func ExampleIndexOfInt16() {

    input := []int16{
        1,
        2,
        3,
        2,
    }
	element := int16(3)
	output := slice.IndexOfInt16(element, input)

    fmt.Printf("%+v", output)

}

// Int32 Example
func ExampleIndexOfInt32() {

    input := []int32{
        1,
        2,
        3,
        2,
    }
	element := int32(3)
	output := slice.IndexOfInt32(element, input)

    fmt.Printf("%+v", output)

}

// Int64 Example
func ExampleIndexOfInt64() {

    input := []int64{
        1,
        2,
        3,
        2,
    }
	element := int64(3)
	output := slice.IndexOfInt64(element, input)

    fmt.Printf("%+v", output)

}

// Uint Example
func ExampleIndexOfUint() {

    input := []uint{
        1,
        2,
        3,
        2,
    }
	element := uint(3)
	output := slice.IndexOfUint(element, input)

    fmt.Printf("%+v", output)

}

// Uint8 Example
func ExampleIndexOfUint8() {

    input := []uint8{
        1,
        2,
        3,
        2,
    }
	element := uint8(3)
	output := slice.IndexOfUint8(element, input)

    fmt.Printf("%+v", output)

}

// Uint16 Example
func ExampleIndexOfUint16() {

    input := []uint16{
        1,
        2,
        3,
        2,
    }
	element := uint16(3)
	output := slice.IndexOfUint16(element, input)

    fmt.Printf("%+v", output)

}

// Uint32 Example
func ExampleIndexOfUint32() {

    input := []uint32{
        1,
        2,
        3,
        2,
    }
	element := uint32(3)
	output := slice.IndexOfUint32(element, input)

    fmt.Printf("%+v", output)

}

// Float32 Example
func ExampleIndexOfFloat32() {

    input := []float32{
        1.1,
        2.2,
        3.3,
        2.2,
    }
	element := float32(3.3)
	output := slice.IndexOfFloat32(element, input)

    fmt.Printf("%+v", output)

}

// Float64 Example
func ExampleIndexOfFloat64() {

    input := []float64{
        1.1,
        2.2,
        3.3,
        2.2,
    }
	element := float64(3.3)
	output := slice.IndexOfFloat64(element, input)

    fmt.Printf("%+v", output)

}

// Complex64 Example
func ExampleIndexOfComplex64() {

    input := []complex64{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.2+3.14i),
        (4.8+3.14i),
    }
	element := complex64((7.2+3.14i))
	output := slice.IndexOfComplex64(element, input)

    fmt.Printf("%+v", output)

}

// Complex128 Example
func ExampleIndexOfComplex128() {

    input := []complex128{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.199999999999999+3.14i),
        (4.8+3.14i),
    }
	element := complex128((7.199999999999999+3.14i))
	output := slice.IndexOfComplex128(element, input)

    fmt.Printf("%+v", output)

}