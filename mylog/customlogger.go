package mylog

import (
	"fmt"
)

type CustomLogger struct{}

func (x CustomLogger) Panicf(msg string, args ...interface{}) {
	fmt.Println("[PANIC]" + msg)
}
func (x CustomLogger) Fatalf(msg string, args ...interface{}) {
	fmt.Println("[FATAL]" + msg)
}
func (x CustomLogger) Errorf(msg string, args ...interface{}) {
	fmt.Println("[ERROR]" + msg)
}
func (x CustomLogger) Warnf(msg string, args ...interface{}) {
	fmt.Println("[WARN]" + msg)
}
func (x CustomLogger) Infof(msg string, args ...interface{}) {
	fmt.Println("[INFO]" + msg)
}
func (x CustomLogger) Debugf(msg string, args ...interface{}) {
	fmt.Println("[DEBUG]" + msg)
}
