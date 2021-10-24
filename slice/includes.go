package slice

// IncludesString func
// Returns true if the input contains the one or more of the values
func IncludesString(input []string, values ...string) bool {
	for _, value := range values {
		idx := IndexOfString(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesInt func
// Returns true if the input contains the one or more of the values
func IncludesInt(input []int, values ...int) bool {
	for _, value := range values {
		idx := IndexOfInt(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesInt8 func
// Returns true if the input contains the one or more of the values
func IncludesInt8(input []int8, values ...int8) bool {
	for _, value := range values {
		idx := IndexOfInt8(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesInt16 func
// Returns true if the input contains the one or more of the values
func IncludesInt16(input []int16, values ...int16) bool {
	for _, value := range values {
		idx := IndexOfInt16(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesInt32 func
// Returns true if the input contains the one or more of the values
func IncludesInt32(input []int32, values ...int32) bool {
	for _, value := range values {
		idx := IndexOfInt32(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesInt64 func
// Returns true if the input contains the one or more of the values
func IncludesInt64(input []int64, values ...int64) bool {
	for _, value := range values {
		idx := IndexOfInt64(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesUint func
// Returns true if the input contains the one or more of the values
func IncludesUint(input []uint, values ...uint) bool {
	for _, value := range values {
		idx := IndexOfUint(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesUint8 func
// Returns true if the input contains the one or more of the values
func IncludesUint8(input []uint8, values ...uint8) bool {
	for _, value := range values {
		idx := IndexOfUint8(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesUint16 func
// Returns true if the input contains the one or more of the values
func IncludesUint16(input []uint16, values ...uint16) bool {
	for _, value := range values {
		idx := IndexOfUint16(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesUint32 func
// Returns true if the input contains the one or more of the values
func IncludesUint32(input []uint32, values ...uint32) bool {
	for _, value := range values {
		idx := IndexOfUint32(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesFloat32 func
// Returns true if the input contains the one or more of the values
func IncludesFloat32(input []float32, values ...float32) bool {
	for _, value := range values {
		idx := IndexOfFloat32(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesFloat64 func
// Returns true if the input contains the one or more of the values
func IncludesFloat64(input []float64, values ...float64) bool {
	for _, value := range values {
		idx := IndexOfFloat64(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesComplex64 func
// Returns true if the input contains the one or more of the values
func IncludesComplex64(input []complex64, values ...complex64) bool {
	for _, value := range values {
		idx := IndexOfComplex64(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}

// IncludesComplex128 func
// Returns true if the input contains the one or more of the values
func IncludesComplex128(input []complex128, values ...complex128) bool {
	for _, value := range values {
		idx := IndexOfComplex128(value, input)
		if idx > -1 {
			return true
		}
	}
	return false
}
