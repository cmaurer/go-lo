package slice


// FilterString func
// Returns a new slice containing the results that the function returns truthy for.
func FilterString(input []string, f func(string) bool) []string {
    var output []string
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterInt func
// Returns a new slice containing the results that the function returns truthy for.
func FilterInt(input []int, f func(int) bool) []int {
    var output []int
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterInt8 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterInt8(input []int8, f func(int8) bool) []int8 {
    var output []int8
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterInt16 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterInt16(input []int16, f func(int16) bool) []int16 {
    var output []int16
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterInt32 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterInt32(input []int32, f func(int32) bool) []int32 {
    var output []int32
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterInt64 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterInt64(input []int64, f func(int64) bool) []int64 {
    var output []int64
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterUint func
// Returns a new slice containing the results that the function returns truthy for.
func FilterUint(input []uint, f func(uint) bool) []uint {
    var output []uint
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterUint8 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterUint8(input []uint8, f func(uint8) bool) []uint8 {
    var output []uint8
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterUint16 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterUint16(input []uint16, f func(uint16) bool) []uint16 {
    var output []uint16
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterUint32 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterUint32(input []uint32, f func(uint32) bool) []uint32 {
    var output []uint32
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterFloat32 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterFloat32(input []float32, f func(float32) bool) []float32 {
    var output []float32
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterFloat64 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterFloat64(input []float64, f func(float64) bool) []float64 {
    var output []float64
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterComplex64 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterComplex64(input []complex64, f func(complex64) bool) []complex64 {
    var output []complex64
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}

// FilterComplex128 func
// Returns a new slice containing the results that the function returns truthy for.
func FilterComplex128(input []complex128, f func(complex128) bool) []complex128 {
    var output []complex128
    for _, value := range input {
        if f(value) {
            output = append(output, value)
        }
    }
    return output
}
