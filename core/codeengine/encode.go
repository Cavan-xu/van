package codeengine



func EncodeBool(buff []byte, b bool) []byte {
	if b {
		return append(buff, byte(1))
	} else {
		return append(buff, byte(0))
	}
}

func EncodeInt8(buff []byte, i int8) []byte {
	return append(buff, byte(i))
}

func EncodeInt16(buff []byte, i int16) []byte {

}

func EncodeInt32(buff []byte, i int32) []byte {

}

func EncodeInt64(buff []byte, i int64) []byte {

}

func EncodeString(buff []byte, s string) []byte {

}