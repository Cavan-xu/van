package log

import "log"

type Log struct {
}

func (l *Log) LogInfo(info string, v ...interface{}) {
	log.Printf(info, v)
}

func (l *Log) LogErr(err error) {
	log.Println(err)
}
