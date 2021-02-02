package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var logLevel string

func main() {
	flag.Parse()
	if logLevel == "" {
		logLevel = "error"
	}
	level, err := logrus.ParseLevel(logLevel)
	if err == nil {
		logrus.SetLevel(level)
	}
	//开始执行
	Run()
}
