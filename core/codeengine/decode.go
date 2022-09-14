package codeengine

import "errors"

func DecodeBool(buff []byte, pos int) (bool, int, error) {

}

func DecodeInt8(buff []byte, pos int) (int8, int, error) {
	if len(buff) <= pos {
		return 0, pos, errors.New("pos is more than buff length")
	}

	return int8(buff[pos]), pos+1, nil
}

func DecodeInt16(buff []byte, pos int) (int16, int, error) {

}

func DecodeInt32(buff []byte, pos int) (int32, int, error) {

}

func DecodeInt64(buff []byte, pos int) (int64, int, error) {

}

func DecodeString(buff []byte, pos int) (string, int, error) {

}