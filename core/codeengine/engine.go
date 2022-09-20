package codeengine

type CodeEngine struct {
	Buff []byte
}

func NewCodeEngine() *CodeEngine {
	return &CodeEngine{
		Buff: make([]byte, 0),
	}
}

func (d *CodeEngine) GetBuff() []byte {
	return d.Buff
}

func (d *CodeEngine) String() string {
	return string(d.Buff)
}
