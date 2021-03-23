package mylog

import (
	"fmt"

	"github.com/hyperledger/aries-framework-go/spi/log"
)

type CustomLoggerProvider struct{}

func NewProvider() log.LoggerProvider {
	return CustomLoggerProvider{}
}

func (x CustomLoggerProvider) GetLogger(module string) log.Logger {
	fmt.Printf("custom logger created.")
	c := CustomLogger{}
	c.Debugf("inside custom logger")
	return c
}
