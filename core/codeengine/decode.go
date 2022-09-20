package codeengine

import "errors"

func (d *CodeEngine) DecodeBool(pos int) (bool, int, error) {
	if pos >= len(d.Buff) {
		return false, pos, errors.New("decode buff length is not enough")
	}

	if d.Buff[pos] == byte(1) {
		return true, pos + 1, nil
	} else {
		return false, pos + 1, nil
	}
}

func (d *CodeEngine) DecodeInt8(pos int) (int8, int, error) {
	if pos > len(d.Buff) {
		return 0, pos, errors.New("decode buff length is not enough")
	}

	return int8(d.Buff[pos]), pos + 1, nil
}

func (d *CodeEngine) DecodeInt16(pos int) (int16, int, error) {
	if pos+2 > len(d.Buff) {
		return 0, pos, errors.New("decode buff length is not enough")
	}

	v1 := int16(d.Buff[pos])
	v2 := int16(d.Buff[pos+1])
	target := v1<<8 | v2
	return target, pos + 2, nil
}

func (d *CodeEngine) DecodeInt32(pos int) (int32, int, error) {
	if pos+4 > len(d.Buff) {
		return 0, pos, errors.New("decode buff length is not enough")
	}

	v1 := int32(d.Buff[pos])
	v2 := int32(d.Buff[pos+1])
	v3 := int32(d.Buff[pos+2])
	v4 := int32(d.Buff[pos+3])
	target := v1<<24 | v2<<16 | v3<<8 | v4
	return target, pos + 4, nil
}

func (d *CodeEngine) DecodeInt64(pos int) (int64, int, error) {
	if pos+8 > len(d.Buff) {
		return 0, pos, errors.New("decode buff length is not enough")
	}

	v1 := int64(d.Buff[pos])
	v2 := int64(d.Buff[pos+1])
	v3 := int64(d.Buff[pos+2])
	v4 := int64(d.Buff[pos+3])
	v5 := int64(d.Buff[pos+4])
	v6 := int64(d.Buff[pos+5])
	v7 := int64(d.Buff[pos+6])
	v8 := int64(d.Buff[pos+7])
	target := v1<<56 | v2<<48 | v3<<40 | v4<<32 | v5<<24 | v6<<16 | v7<<8 | v8
	return target, pos + 8, nil
}

func (d *CodeEngine) DecodeString(pos int) (string, int, error) {
	strLen, pos, err := d.DecodeInt16(pos)
	if err != nil {
		return "", pos, err
	}

	if pos+int(strLen) > len(d.Buff) {
		return "", pos, errors.New("decode buff length is not enough")
	}

	tmp := make([]byte, strLen)
	copy(tmp, d.Buff[pos:])
	return string(tmp), pos + int(strLen), nil
}
