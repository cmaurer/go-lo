package slice


// FindString func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindString(input []string, f func(string) bool) *string {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindInt func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindInt(input []int, f func(int) bool) *int {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindInt8 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindInt8(input []int8, f func(int8) bool) *int8 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindInt16 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindInt16(input []int16, f func(int16) bool) *int16 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindInt32 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindInt32(input []int32, f func(int32) bool) *int32 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindInt64 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindInt64(input []int64, f func(int64) bool) *int64 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindUint func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindUint(input []uint, f func(uint) bool) *uint {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindUint8 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindUint8(input []uint8, f func(uint8) bool) *uint8 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindUint16 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindUint16(input []uint16, f func(uint16) bool) *uint16 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindUint32 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindUint32(input []uint32, f func(uint32) bool) *uint32 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindFloat32 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindFloat32(input []float32, f func(float32) bool) *float32 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindFloat64 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindFloat64(input []float64, f func(float64) bool) *float64 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindComplex64 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindComplex64(input []complex64, f func(complex64) bool) *complex64 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}

// FindComplex128 func
// Returns a pointer to the first element that the function returns truthy for, otherwise returns nil if the element is not found.
func FindComplex128(input []complex128, f func(complex128) bool) *complex128 {
    for _, value := range input {
        if f(value) {
            return &value
        }
    }
    return nil
}
