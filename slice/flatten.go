package slice


// FlattenString func
// Returns a new concatenated slice from one or more provided slices.
func FlattenString(input ...[]string) []string {
    var output []string
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenInt func
// Returns a new concatenated slice from one or more provided slices.
func FlattenInt(input ...[]int) []int {
    var output []int
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenInt8 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenInt8(input ...[]int8) []int8 {
    var output []int8
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenInt16 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenInt16(input ...[]int16) []int16 {
    var output []int16
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenInt32 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenInt32(input ...[]int32) []int32 {
    var output []int32
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenInt64 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenInt64(input ...[]int64) []int64 {
    var output []int64
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenUint func
// Returns a new concatenated slice from one or more provided slices.
func FlattenUint(input ...[]uint) []uint {
    var output []uint
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenUint8 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenUint8(input ...[]uint8) []uint8 {
    var output []uint8
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenUint16 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenUint16(input ...[]uint16) []uint16 {
    var output []uint16
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenUint32 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenUint32(input ...[]uint32) []uint32 {
    var output []uint32
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenFloat32 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenFloat32(input ...[]float32) []float32 {
    var output []float32
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenFloat64 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenFloat64(input ...[]float64) []float64 {
    var output []float64
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenComplex64 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenComplex64(input ...[]complex64) []complex64 {
    var output []complex64
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}

// FlattenComplex128 func
// Returns a new concatenated slice from one or more provided slices.
func FlattenComplex128(input ...[]complex128) []complex128 {
    var output []complex128
    for _, in := range input {
        for _, inn := range in {
            output = append(output, inn)
        }
    }
    return output
}
