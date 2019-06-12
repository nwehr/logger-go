package logger

import (
	"encoding/json"
	"os"
	"runtime"
	"time"
)

type LogItem struct {
	Severity  Severity    `json:"severity"`
	Timestamp time.Time   `json:"timestamp"`
	Message   interface{} `json:"message"`
	Tags      []Tag       `json:"tags"`
	Caller    Caller      `json:"caller"`
}

type Tag string
type Severity string

const (
	DEBUG   = "debug"
	WARNING = "warning"
	ERROR   = "error"
	FATAL   = "fatal"
)

type Caller struct {
	File string `json:"file"`
	Line int    `json:"line"`
	Func string `json:"func"`
}

func Debug(message interface{}, tags ...Tag) {
	logItem := makeLogItem(message, tags...)
	printLogItem(logItem)
}

func Warning(message interface{}, tags ...Tag) {
	logItem := makeLogItem(message, tags...)
	logItem.Severity = WARNING

	printLogItem(logItem)
}

func Error(message interface{}, tags ...Tag) {
	logItem := makeLogItem(message, tags...)
	logItem.Severity = ERROR

	printLogItem(logItem)
}

func Fatal(message interface{}, tags ...Tag) {
	logItem := makeLogItem(message, tags...)
	logItem.Severity = FATAL

	printLogItem(logItem)
	os.Exit(1)
}

func makeLogItem(message interface{}, tags ...Tag) LogItem {
	if tags == nil {
		tags = []Tag{}
	}

	return LogItem{
		Severity:  DEBUG,
		Timestamp: time.Now(),
		Message:   message,
		Tags:      tags,
		Caller:    caller(3),
	}
}

func caller(depth int) Caller {
	programCounter, file, line, _ := runtime.Caller(depth)
	funcDetails := runtime.FuncForPC(programCounter)

	return Caller{
		File: file,
		Line: line,
		Func: funcDetails.Name(),
	}
}

func printLogItem(logItem LogItem) {
	data, _ := json.Marshal(logItem)
	println(string(data))
}
