package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

var logLevel string

func main() {
	flag.Parse()
	logLevel = os.Getenv("GIT_RELEASENOTE_LOG_LEVEL")
	if logLevel == "" {
		logLevel = "error"
	}
	level, err := logrus.ParseLevel(logLevel)
	logrus.SetOutput(os.Stderr)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if err == nil {
		logrus.SetLevel(level)
	}
	//开始执行
	Run()
}
