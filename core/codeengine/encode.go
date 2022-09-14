package codeengine

func (d *CodeEngine) EncodeBool(b bool) {
	if b {
		d.Buff = append(d.Buff, byte(1))
	} else {
		d.Buff = append(d.Buff, byte(0))
	}
}

func (d *CodeEngine) EncodeInt8(i int8) {
	d.Buff = append(d.Buff, byte(i))
}

func (d *CodeEngine) EncodeInt16(i int16) {
	d.Buff = append(d.Buff, byte(i>>8), byte(i))
}

func (d *CodeEngine) EncodeInt32(i int32) {
	d.Buff = append(d.Buff, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

func (d *CodeEngine) EncodeInt64(i int64) {
	d.Buff = append(d.Buff, byte(i>>56), byte(i>>48), byte(i>>40), byte(i>>32), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

func (d *CodeEngine) EncodeString(s string) {
	d.EncodeInt16(int16(len(s)))
	d.Buff = append(d.Buff, []byte(s)...)
}
