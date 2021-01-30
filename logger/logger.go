package logger

import (
	"github.com/favframework/debug"
	"log"
)

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Info 输出信息
func Info(inter interface{}) {
	godump.Dump(inter)
}