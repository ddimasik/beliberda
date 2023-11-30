package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
	"fmt"
)

func main() {
	var delay int
	var size int
	var help bool
	flag.IntVar(&delay, "delay", 3, "the delay in seconds between sending messages")
	flag.IntVar(&size, "size", 1, "the size of the log message in kilobytes")
    flag.BoolVar(&help, "help", false, "display help")

    flag.Parse()

    if help {
        fmt.Println("Usage of our program: docker run ddimasik/beliberda -delay 5 -size 2")
        flag.PrintDefaults()
        os.Exit(0)
    }

	textLogger := logrus.New()
	jsonLogger := logrus.New()

	jsonLogger.SetFormatter(&logrus.JSONFormatter{})
	jsonLogger.SetOutput(os.Stdout)

	textLogger.SetFormatter(&logrus.TextFormatter{})
	textLogger.SetOutput(os.Stdout)

	// 1 kilobyte is approximately 1024 characters
	phrase := "This is a meaningful log message. "
	message := strings.Repeat(phrase, (size*1024)/len(phrase))

	for {
		textLogger.Info(message)
		jsonLogger.Info(message)
		time.Sleep(time.Duration(delay) * time.Second)

		textLogger.Warn(message)
		jsonLogger.Warn(message)
		time.Sleep(time.Duration(delay) * time.Second)

		textLogger.Error(message)
		jsonLogger.Error(message)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}