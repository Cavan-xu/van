package log

import "log"

type Log struct {
}

func (l *Log) LogInfo(info string, v ...interface{}) {
	log.Printf(info, v)
}

func (l *Log) LogErr(info string, v ...interface{}) {
	log.Printf(info, v)
}

func (l *Log) LogPanic(info string, v ...interface{}) {
	log.Println(info, v)
}
