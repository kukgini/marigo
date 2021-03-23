package mylog

import (
	"fmt"
)

type CustomLogger struct{}

func (x CustomLogger) Panicf(msg string, args ...interface{}) {
	fmt.Println("PANIC %s", msg)
}
func (x CustomLogger) Fatalf(msg string, args ...interface{}) {
	fmt.Println("PANIC %s", msg)
}
func (x CustomLogger) Errorf(msg string, args ...interface{}) {
	fmt.Println("PANIC %s", msg)
}
func (x CustomLogger) Warnf(msg string, args ...interface{}) {
	fmt.Println("PANIC %s", msg)
}
func (x CustomLogger) Infof(msg string, args ...interface{}) {
	fmt.Println("PANIC %s", msg)
}
func (x CustomLogger) Debugf(msg string, args ...interface{}) {
	fmt.Println("PANIC %s", msg)
}
