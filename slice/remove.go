package slice

// RemoveString func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveString(input []string, f func(string) bool) []string {
	output := []string{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveInt func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveInt(input []int, f func(int) bool) []int {
	output := []int{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveInt8 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveInt8(input []int8, f func(int8) bool) []int8 {
	output := []int8{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveInt16 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveInt16(input []int16, f func(int16) bool) []int16 {
	output := []int16{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveInt32 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveInt32(input []int32, f func(int32) bool) []int32 {
	output := []int32{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveInt64 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveInt64(input []int64, f func(int64) bool) []int64 {
	output := []int64{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveUint func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveUint(input []uint, f func(uint) bool) []uint {
	output := []uint{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveUint8 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveUint8(input []uint8, f func(uint8) bool) []uint8 {
	output := []uint8{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveUint16 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveUint16(input []uint16, f func(uint16) bool) []uint16 {
	output := []uint16{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveUint32 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveUint32(input []uint32, f func(uint32) bool) []uint32 {
	output := []uint32{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveFloat32 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveFloat32(input []float32, f func(float32) bool) []float32 {
	output := []float32{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveFloat64 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveFloat64(input []float64, f func(float64) bool) []float64 {
	output := []float64{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveComplex64 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveComplex64(input []complex64, f func(complex64) bool) []complex64 {
	output := []complex64{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}

// RemoveComplex128 func
// Removes each element that the function returns truthy for and returns a slice with the the items removed.
func RemoveComplex128(input []complex128, f func(complex128) bool) []complex128 {
	output := []complex128{}
	for _, value := range input {
		if !f(value) {
			output = append(output, value)
		}
	}
	return output
}
