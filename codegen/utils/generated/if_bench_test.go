package utils

import (
	"testing"
)
		
func BenchmarkTestIfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfString(true, "p1", "p2")
			IfString(false, "p1", "p2")
		
	}
}

func BenchmarkTestIfStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := "p1"
			IfStringPtr(true, &p1, nil)
			IfStringPtr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfInt(true, int(1), int(2))
			IfInt(false, int(1), int(2))
		
	}
}

func BenchmarkTestIfIntPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := int(1)
			IfIntPtr(true, &p1, nil)
			IfIntPtr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfInt8(true, int8(2), int8(3))
			IfInt8(false, int8(2), int8(3))
		
	}
}

func BenchmarkTestIfInt8Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := int8(2)
			IfInt8Ptr(true, &p1, nil)
			IfInt8Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfInt16(true, int16(4), int16(5))
			IfInt16(false, int16(4), int16(5))
		
	}
}

func BenchmarkTestIfInt16Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := int16(4)
			IfInt16Ptr(true, &p1, nil)
			IfInt16Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfInt32(true, int32(6), int32(7))
			IfInt32(false, int32(6), int32(7))
		
	}
}

func BenchmarkTestIfInt32Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := int32(6)
			IfInt32Ptr(true, &p1, nil)
			IfInt32Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfInt64(true, int64(10), int64(11))
			IfInt64(false, int64(10), int64(11))
		
	}
}

func BenchmarkTestIfInt64Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := int64(10)
			IfInt64Ptr(true, &p1, nil)
			IfInt64Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfUint(true, uint(1), uint(2))
			IfUint(false, uint(1), uint(2))
		
	}
}

func BenchmarkTestIfUintPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := uint(1)
			IfUintPtr(true, &p1, nil)
			IfUintPtr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfUint8(true, uint8(2), uint8(3))
			IfUint8(false, uint8(2), uint8(3))
		
	}
}

func BenchmarkTestIfUint8Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := uint8(2)
			IfUint8Ptr(true, &p1, nil)
			IfUint8Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfUint16(true, uint16(4), uint16(5))
			IfUint16(false, uint16(4), uint16(5))
		
	}
}

func BenchmarkTestIfUint16Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := uint16(4)
			IfUint16Ptr(true, &p1, nil)
			IfUint16Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfUint32(true, uint32(6), uint32(7))
			IfUint32(false, uint32(6), uint32(7))
		
	}
}

func BenchmarkTestIfUint32Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := uint32(6)
			IfUint32Ptr(true, &p1, nil)
			IfUint32Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfFloat32(true, float32(8), float32(7))
			IfFloat32(false, float32(8), float32(7))
		
	}
}

func BenchmarkTestIfFloat32Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := float32(8)
			IfFloat32Ptr(true, &p1, nil)
			IfFloat32Ptr(false, &p1, nil)
		
	}
}

func BenchmarkTestIfFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			IfFloat64(true, float64(9), float64(8))
			IfFloat64(false, float64(9), float64(8))
		
	}
}

func BenchmarkTestIfFloat64Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
			p1 := float64(9)
			IfFloat64Ptr(true, &p1, nil)
			IfFloat64Ptr(false, &p1, nil)
		
	}
}
