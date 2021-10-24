package slice_test

import (
	"fmt"
	"github.com/cmaurer/go-lo/slice"
)

		
// String Example
func ExampleFlattenString() {

    input := []string{
        "test1",
        "test2",
        "test3",
        "test2",
    }

    input2 := []string{
        "test1",
        "test2",
        "test3",
    }

    output := slice.FlattenString(input, input2)

    fmt.Printf("%+v", output)

}
// Int Example
func ExampleFlattenInt() {

    input := []int{
        1,
        2,
        3,
        2,
    }

    input2 := []int{
        1,
        2,
        3,
    }

    output := slice.FlattenInt(input, input2)

    fmt.Printf("%+v", output)

}
// Int8 Example
func ExampleFlattenInt8() {

    input := []int8{
        1,
        2,
        3,
        2,
    }

    input2 := []int8{
        1,
        2,
        3,
    }

    output := slice.FlattenInt8(input, input2)

    fmt.Printf("%+v", output)

}
// Int16 Example
func ExampleFlattenInt16() {

    input := []int16{
        1,
        2,
        3,
        2,
    }

    input2 := []int16{
        1,
        2,
        3,
    }

    output := slice.FlattenInt16(input, input2)

    fmt.Printf("%+v", output)

}
// Int32 Example
func ExampleFlattenInt32() {

    input := []int32{
        1,
        2,
        3,
        2,
    }

    input2 := []int32{
        1,
        2,
        3,
    }

    output := slice.FlattenInt32(input, input2)

    fmt.Printf("%+v", output)

}
// Int64 Example
func ExampleFlattenInt64() {

    input := []int64{
        1,
        2,
        3,
        2,
    }

    input2 := []int64{
        1,
        2,
        3,
    }

    output := slice.FlattenInt64(input, input2)

    fmt.Printf("%+v", output)

}
// Uint Example
func ExampleFlattenUint() {

    input := []uint{
        1,
        2,
        3,
        2,
    }

    input2 := []uint{
        1,
        2,
        3,
    }

    output := slice.FlattenUint(input, input2)

    fmt.Printf("%+v", output)

}
// Uint8 Example
func ExampleFlattenUint8() {

    input := []uint8{
        1,
        2,
        3,
        2,
    }

    input2 := []uint8{
        1,
        2,
        3,
    }

    output := slice.FlattenUint8(input, input2)

    fmt.Printf("%+v", output)

}
// Uint16 Example
func ExampleFlattenUint16() {

    input := []uint16{
        1,
        2,
        3,
        2,
    }

    input2 := []uint16{
        1,
        2,
        3,
    }

    output := slice.FlattenUint16(input, input2)

    fmt.Printf("%+v", output)

}
// Uint32 Example
func ExampleFlattenUint32() {

    input := []uint32{
        1,
        2,
        3,
        2,
    }

    input2 := []uint32{
        1,
        2,
        3,
    }

    output := slice.FlattenUint32(input, input2)

    fmt.Printf("%+v", output)

}
// Float32 Example
func ExampleFlattenFloat32() {

    input := []float32{
        1.1,
        2.2,
        3.3,
        2.2,
    }

    input2 := []float32{
        1.1,
        2.2,
        3.3,
    }

    output := slice.FlattenFloat32(input, input2)

    fmt.Printf("%+v", output)

}
// Float64 Example
func ExampleFlattenFloat64() {

    input := []float64{
        1.1,
        2.2,
        3.3,
        2.2,
    }

    input2 := []float64{
        1.1,
        2.2,
        3.3,
    }

    output := slice.FlattenFloat64(input, input2)

    fmt.Printf("%+v", output)

}
// Complex64 Example
func ExampleFlattenComplex64() {

    input := []complex64{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.2+3.14i),
        (4.8+3.14i),
    }

    input2 := []complex64{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.2+3.14i),
    }

    output := slice.FlattenComplex64(input, input2)

    fmt.Printf("%+v", output)

}
// Complex128 Example
func ExampleFlattenComplex128() {

    input := []complex128{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.199999999999999+3.14i),
        (4.8+3.14i),
    }

    input2 := []complex128{
        (2.4+3.14i),
        (4.8+3.14i),
        (7.199999999999999+3.14i),
    }

    output := slice.FlattenComplex128(input, input2)

    fmt.Printf("%+v", output)

}