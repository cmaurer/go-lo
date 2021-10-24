package slice


// ReduceString func
// Returns a string value by executing the accumulator function for each input value.
func ReduceString(input []string, f func(string, string) string, acc string) string {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceInt func
// Returns a int value by executing the accumulator function for each input value.
func ReduceInt(input []int, f func(int, int) int, acc int) int {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceInt8 func
// Returns a int8 value by executing the accumulator function for each input value.
func ReduceInt8(input []int8, f func(int8, int8) int8, acc int8) int8 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceInt16 func
// Returns a int16 value by executing the accumulator function for each input value.
func ReduceInt16(input []int16, f func(int16, int16) int16, acc int16) int16 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceInt32 func
// Returns a int32 value by executing the accumulator function for each input value.
func ReduceInt32(input []int32, f func(int32, int32) int32, acc int32) int32 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceInt64 func
// Returns a int64 value by executing the accumulator function for each input value.
func ReduceInt64(input []int64, f func(int64, int64) int64, acc int64) int64 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceUint func
// Returns a uint value by executing the accumulator function for each input value.
func ReduceUint(input []uint, f func(uint, uint) uint, acc uint) uint {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceUint8 func
// Returns a uint8 value by executing the accumulator function for each input value.
func ReduceUint8(input []uint8, f func(uint8, uint8) uint8, acc uint8) uint8 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceUint16 func
// Returns a uint16 value by executing the accumulator function for each input value.
func ReduceUint16(input []uint16, f func(uint16, uint16) uint16, acc uint16) uint16 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceUint32 func
// Returns a uint32 value by executing the accumulator function for each input value.
func ReduceUint32(input []uint32, f func(uint32, uint32) uint32, acc uint32) uint32 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceFloat32 func
// Returns a float32 value by executing the accumulator function for each input value.
func ReduceFloat32(input []float32, f func(float32, float32) float32, acc float32) float32 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceFloat64 func
// Returns a float64 value by executing the accumulator function for each input value.
func ReduceFloat64(input []float64, f func(float64, float64) float64, acc float64) float64 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceComplex64 func
// Returns a complex64 value by executing the accumulator function for each input value.
func ReduceComplex64(input []complex64, f func(complex64, complex64) complex64, acc complex64) complex64 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}

// ReduceComplex128 func
// Returns a complex128 value by executing the accumulator function for each input value.
func ReduceComplex128(input []complex128, f func(complex128, complex128) complex128, acc complex128) complex128 {
	for _, s := range input {
	   acc = f(acc, s)
	}
	return acc
}
