package codeengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeEngine_Bool(t *testing.T) {
	source := false

	engine := NewCodeEngine()
	engine.EncodeBool(source)
	target, pos, err := engine.DecodeBool(0)

	assert.Equal(t, source, target)
	assert.Equal(t, 1, pos)
	assert.Nil(t, err)
}

func TestCodeEngine_int8(t *testing.T) {
	var source int8 = 10

	engine := NewCodeEngine()
	engine.EncodeInt8(source)
	target, pos, err := engine.DecodeInt8(0)

	assert.Equal(t, source, target)
	assert.Equal(t, 1, pos)
	assert.Nil(t, err)
}

func TestCodeEngine_int16(t *testing.T) {
	var source int16 = 10

	engine := NewCodeEngine()
	engine.EncodeInt16(source)
	target, pos, err := engine.DecodeInt16(0)

	assert.Equal(t, source, target)
	assert.Equal(t, 2, pos)
	assert.Nil(t, err)
}

func TestCodeEngine_int32(t *testing.T) {
	var source int32 = 10

	engine := NewCodeEngine()
	engine.EncodeInt32(source)
	target, pos, err := engine.DecodeInt32(0)

	assert.Equal(t, source, target)
	assert.Equal(t, 4, pos)
	assert.Nil(t, err)
}

func TestCodeEngine_int64(t *testing.T) {
	var source int64 = 10

	engine := NewCodeEngine()
	engine.EncodeInt64(source)
	target, pos, err := engine.DecodeInt64(0)

	assert.Equal(t, source, target)
	assert.Equal(t, 8, pos)
	assert.Nil(t, err)
}

func TestCodeEngine_string(t *testing.T) {
	source := "hello world"

	engine := NewCodeEngine()
	engine.EncodeString(source)
	target, pos, err := engine.DecodeString(0)

	assert.Equal(t, source, target)
	assert.Equal(t, 13, pos)
	assert.Nil(t, err)
}

func TestCodeEngine_use(t *testing.T) {
	engine := NewCodeEngine()
	engine.EncodeBool(false)
	engine.EncodeInt64(1)
	engine.EncodeInt16(8)
	engine.EncodeString("hello")

	val1, pos, err := engine.DecodeBool(0)
	assert.Equal(t, false, val1)
	assert.Nil(t, err)

	val2, pos, err := engine.DecodeInt64(pos)
	assert.Equal(t, int64(1), val2)
	assert.Nil(t, err)

	val3, pos, err := engine.DecodeInt16(pos)
	assert.Equal(t, int16(8), val3)
	assert.Nil(t, err)

	val4, pos, err := engine.DecodeString(pos)
	assert.Equal(t, "hello", val4)
	assert.Nil(t, err)
}
