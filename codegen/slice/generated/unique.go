package slice


// UniqueString func
// Returns a slice without duplicates
func UniqueString(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueInt func
// Returns a slice without duplicates
func UniqueInt(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueInt8 func
// Returns a slice without duplicates
func UniqueInt8(input []int8) []int8 {
	u := make([]int8, 0, len(input))
	m := make(map[int8]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueInt16 func
// Returns a slice without duplicates
func UniqueInt16(input []int16) []int16 {
	u := make([]int16, 0, len(input))
	m := make(map[int16]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueInt32 func
// Returns a slice without duplicates
func UniqueInt32(input []int32) []int32 {
	u := make([]int32, 0, len(input))
	m := make(map[int32]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueInt64 func
// Returns a slice without duplicates
func UniqueInt64(input []int64) []int64 {
	u := make([]int64, 0, len(input))
	m := make(map[int64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueUint func
// Returns a slice without duplicates
func UniqueUint(input []uint) []uint {
	u := make([]uint, 0, len(input))
	m := make(map[uint]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueUint8 func
// Returns a slice without duplicates
func UniqueUint8(input []uint8) []uint8 {
	u := make([]uint8, 0, len(input))
	m := make(map[uint8]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueUint16 func
// Returns a slice without duplicates
func UniqueUint16(input []uint16) []uint16 {
	u := make([]uint16, 0, len(input))
	m := make(map[uint16]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueUint32 func
// Returns a slice without duplicates
func UniqueUint32(input []uint32) []uint32 {
	u := make([]uint32, 0, len(input))
	m := make(map[uint32]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueFloat32 func
// Returns a slice without duplicates
func UniqueFloat32(input []float32) []float32 {
	u := make([]float32, 0, len(input))
	m := make(map[float32]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueFloat64 func
// Returns a slice without duplicates
func UniqueFloat64(input []float64) []float64 {
	u := make([]float64, 0, len(input))
	m := make(map[float64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueComplex64 func
// Returns a slice without duplicates
func UniqueComplex64(input []complex64) []complex64 {
	u := make([]complex64, 0, len(input))
	m := make(map[complex64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

// UniqueComplex128 func
// Returns a slice without duplicates
func UniqueComplex128(input []complex128) []complex128 {
	u := make([]complex128, 0, len(input))
	m := make(map[complex128]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
