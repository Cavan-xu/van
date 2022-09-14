package codeengine

type CodeEngine struct {
	Buff []byte
}

func NewCodeEngine() *CodeEngine {
	return &CodeEngine{
		Buff: make([]byte, 0),
	}
}
