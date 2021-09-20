package utils

// IfString func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfString(true, "p1", "p2")
func IfString(b bool, p1 string, p2 string) string {
	if b {
		return p1
	}
	return p2
}

// IfStringPtr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfStringPtr(true, "p1", nil)
func IfStringPtr(b bool, p1 *string, p2 *string) *string {
	if b {
		return p1
	}
	return p2
}

// IfInt func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt(true, int(1), int(2))
func IfInt(b bool, p1 int, p2 int) int {
	if b {
		return p1
	}
	return p2
}

// IfIntPtr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfIntPtr(true, int(1), nil)
func IfIntPtr(b bool, p1 *int, p2 *int) *int {
	if b {
		return p1
	}
	return p2
}

// IfInt8 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt8(true, int8(2), int8(3))
func IfInt8(b bool, p1 int8, p2 int8) int8 {
	if b {
		return p1
	}
	return p2
}

// IfInt8Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt8Ptr(true, int8(2), nil)
func IfInt8Ptr(b bool, p1 *int8, p2 *int8) *int8 {
	if b {
		return p1
	}
	return p2
}

// IfInt16 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt16(true, int16(4), int16(5))
func IfInt16(b bool, p1 int16, p2 int16) int16 {
	if b {
		return p1
	}
	return p2
}

// IfInt16Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt16Ptr(true, int16(4), nil)
func IfInt16Ptr(b bool, p1 *int16, p2 *int16) *int16 {
	if b {
		return p1
	}
	return p2
}

// IfInt32 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt32(true, int32(6), int32(7))
func IfInt32(b bool, p1 int32, p2 int32) int32 {
	if b {
		return p1
	}
	return p2
}

// IfInt32Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt32Ptr(true, int32(6), nil)
func IfInt32Ptr(b bool, p1 *int32, p2 *int32) *int32 {
	if b {
		return p1
	}
	return p2
}

// IfInt64 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt64(true, int64(10), int64(11))
func IfInt64(b bool, p1 int64, p2 int64) int64 {
	if b {
		return p1
	}
	return p2
}

// IfInt64Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfInt64Ptr(true, int64(10), nil)
func IfInt64Ptr(b bool, p1 *int64, p2 *int64) *int64 {
	if b {
		return p1
	}
	return p2
}

// IfUint func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUint(true, uint(1), uint(2))
func IfUint(b bool, p1 uint, p2 uint) uint {
	if b {
		return p1
	}
	return p2
}

// IfUintPtr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUintPtr(true, uint(1), nil)
func IfUintPtr(b bool, p1 *uint, p2 *uint) *uint {
	if b {
		return p1
	}
	return p2
}

// IfUint8 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUint8(true, uint8(2), uint8(3))
func IfUint8(b bool, p1 uint8, p2 uint8) uint8 {
	if b {
		return p1
	}
	return p2
}

// IfUint8Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUint8Ptr(true, uint8(2), nil)
func IfUint8Ptr(b bool, p1 *uint8, p2 *uint8) *uint8 {
	if b {
		return p1
	}
	return p2
}

// IfUint16 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUint16(true, uint16(4), uint16(5))
func IfUint16(b bool, p1 uint16, p2 uint16) uint16 {
	if b {
		return p1
	}
	return p2
}

// IfUint16Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUint16Ptr(true, uint16(4), nil)
func IfUint16Ptr(b bool, p1 *uint16, p2 *uint16) *uint16 {
	if b {
		return p1
	}
	return p2
}

// IfUint32 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUint32(true, uint32(6), uint32(7))
func IfUint32(b bool, p1 uint32, p2 uint32) uint32 {
	if b {
		return p1
	}
	return p2
}

// IfUint32Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfUint32Ptr(true, uint32(6), nil)
func IfUint32Ptr(b bool, p1 *uint32, p2 *uint32) *uint32 {
	if b {
		return p1
	}
	return p2
}

// IfFloat32 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfFloat32(true, float32(8), float32(7))
func IfFloat32(b bool, p1 float32, p2 float32) float32 {
	if b {
		return p1
	}
	return p2
}

// IfFloat32Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfFloat32Ptr(true, float32(8), nil)
func IfFloat32Ptr(b bool, p1 *float32, p2 *float32) *float32 {
	if b {
		return p1
	}
	return p2
}

// IfFloat64 func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfFloat64(true, float64(9), float64(8))
func IfFloat64(b bool, p1 float64, p2 float64) float64 {
	if b {
		return p1
	}
	return p2
}

// IfFloat64Ptr func
// If the `b` parameter is true, then the p1 value will be returned, otherwise the p2 value will be returned.
// for example:
//    IfFloat64Ptr(true, float64(9), nil)
func IfFloat64Ptr(b bool, p1 *float64, p2 *float64) *float64 {
	if b {
		return p1
	}
	return p2
}
