package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfString(t *testing.T) {

	assert.Equal(t, IfString(true, "p1", "p2"), "p1")
	assert.Equal(t, IfString(false, "p1", "p2"), "p2")

}

func TestIfStringPtr(t *testing.T) {

	p1 := "p1"

	assert.Equal(t, IfStringPtr(true, &p1, nil), &p1)
	assert.Nil(t, IfStringPtr(false, &p1, nil))

}

func TestIfInt(t *testing.T) {

	assert.Equal(t, IfInt(true, int(1), int(2)), int(1))
	assert.Equal(t, IfInt(false, int(1), int(2)), int(2))

}

func TestIfIntPtr(t *testing.T) {

	p1 := int(1)

	assert.Equal(t, IfIntPtr(true, &p1, nil), &p1)
	assert.Nil(t, IfIntPtr(false, &p1, nil))

}

func TestIfInt8(t *testing.T) {

	assert.Equal(t, IfInt8(true, int8(2), int8(3)), int8(2))
	assert.Equal(t, IfInt8(false, int8(2), int8(3)), int8(3))

}

func TestIfInt8Ptr(t *testing.T) {

	p1 := int8(2)

	assert.Equal(t, IfInt8Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfInt8Ptr(false, &p1, nil))

}

func TestIfInt16(t *testing.T) {

	assert.Equal(t, IfInt16(true, int16(4), int16(5)), int16(4))
	assert.Equal(t, IfInt16(false, int16(4), int16(5)), int16(5))

}

func TestIfInt16Ptr(t *testing.T) {

	p1 := int16(4)

	assert.Equal(t, IfInt16Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfInt16Ptr(false, &p1, nil))

}

func TestIfInt32(t *testing.T) {

	assert.Equal(t, IfInt32(true, int32(6), int32(7)), int32(6))
	assert.Equal(t, IfInt32(false, int32(6), int32(7)), int32(7))

}

func TestIfInt32Ptr(t *testing.T) {

	p1 := int32(6)

	assert.Equal(t, IfInt32Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfInt32Ptr(false, &p1, nil))

}

func TestIfInt64(t *testing.T) {

	assert.Equal(t, IfInt64(true, int64(10), int64(11)), int64(10))
	assert.Equal(t, IfInt64(false, int64(10), int64(11)), int64(11))

}

func TestIfInt64Ptr(t *testing.T) {

	p1 := int64(10)

	assert.Equal(t, IfInt64Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfInt64Ptr(false, &p1, nil))

}

func TestIfUint(t *testing.T) {

	assert.Equal(t, IfUint(true, uint(1), uint(2)), uint(1))
	assert.Equal(t, IfUint(false, uint(1), uint(2)), uint(2))

}

func TestIfUintPtr(t *testing.T) {

	p1 := uint(1)

	assert.Equal(t, IfUintPtr(true, &p1, nil), &p1)
	assert.Nil(t, IfUintPtr(false, &p1, nil))

}

func TestIfUint8(t *testing.T) {

	assert.Equal(t, IfUint8(true, uint8(2), uint8(3)), uint8(2))
	assert.Equal(t, IfUint8(false, uint8(2), uint8(3)), uint8(3))

}

func TestIfUint8Ptr(t *testing.T) {

	p1 := uint8(2)

	assert.Equal(t, IfUint8Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfUint8Ptr(false, &p1, nil))

}

func TestIfUint16(t *testing.T) {

	assert.Equal(t, IfUint16(true, uint16(4), uint16(5)), uint16(4))
	assert.Equal(t, IfUint16(false, uint16(4), uint16(5)), uint16(5))

}

func TestIfUint16Ptr(t *testing.T) {

	p1 := uint16(4)

	assert.Equal(t, IfUint16Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfUint16Ptr(false, &p1, nil))

}

func TestIfUint32(t *testing.T) {

	assert.Equal(t, IfUint32(true, uint32(6), uint32(7)), uint32(6))
	assert.Equal(t, IfUint32(false, uint32(6), uint32(7)), uint32(7))

}

func TestIfUint32Ptr(t *testing.T) {

	p1 := uint32(6)

	assert.Equal(t, IfUint32Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfUint32Ptr(false, &p1, nil))

}

func TestIfFloat32(t *testing.T) {

	assert.Equal(t, IfFloat32(true, float32(8), float32(7)), float32(8))
	assert.Equal(t, IfFloat32(false, float32(8), float32(7)), float32(7))

}

func TestIfFloat32Ptr(t *testing.T) {

	p1 := float32(8)

	assert.Equal(t, IfFloat32Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfFloat32Ptr(false, &p1, nil))

}

func TestIfFloat64(t *testing.T) {

	assert.Equal(t, IfFloat64(true, float64(9), float64(8)), float64(9))
	assert.Equal(t, IfFloat64(false, float64(9), float64(8)), float64(8))

}

func TestIfFloat64Ptr(t *testing.T) {

	p1 := float64(9)

	assert.Equal(t, IfFloat64Ptr(true, &p1, nil), &p1)
	assert.Nil(t, IfFloat64Ptr(false, &p1, nil))

}
