package main

import (
	"log"

	"github.com/fatih/color"
)

type Logger struct{}

func (l *Logger) Verbose(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *Logger) Warning(format string, v ...interface{}) {
	yellow := color.New(color.FgYellow).SprintFunc()
	log.Printf(yellow(format), v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	log.Printf(red(format), v...)
}
