package slice


// MapString func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapString(input []string, f func(string) string) []string {
    inputMap := make([]string, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapInt func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapInt(input []int, f func(int) int) []int {
    inputMap := make([]int, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapInt8 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapInt8(input []int8, f func(int8) int8) []int8 {
    inputMap := make([]int8, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapInt16 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapInt16(input []int16, f func(int16) int16) []int16 {
    inputMap := make([]int16, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapInt32 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapInt32(input []int32, f func(int32) int32) []int32 {
    inputMap := make([]int32, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapInt64 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapInt64(input []int64, f func(int64) int64) []int64 {
    inputMap := make([]int64, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapUint func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapUint(input []uint, f func(uint) uint) []uint {
    inputMap := make([]uint, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapUint8 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapUint8(input []uint8, f func(uint8) uint8) []uint8 {
    inputMap := make([]uint8, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapUint16 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapUint16(input []uint16, f func(uint16) uint16) []uint16 {
    inputMap := make([]uint16, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapUint32 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapUint32(input []uint32, f func(uint32) uint32) []uint32 {
    inputMap := make([]uint32, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapFloat32 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapFloat32(input []float32, f func(float32) float32) []float32 {
    inputMap := make([]float32, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapFloat64 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapFloat64(input []float64, f func(float64) float64) []float64 {
    inputMap := make([]float64, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapComplex64 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapComplex64(input []complex64, f func(complex64) complex64) []complex64 {
    inputMap := make([]complex64, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}

// MapComplex128 func
// Returns a new slice containing the results from applying the func to each element in the slice.
func MapComplex128(input []complex128, f func(complex128) complex128) []complex128 {
    inputMap := make([]complex128, len(input))
    for i, value := range input {
        inputMap[i] = f(value)
    }
    return inputMap
}
