package logger

import (
	"bytes"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	expected := sevDebug
	got := debug("Hello").Severity

	if expected != got {
		t.Errorf("expected %s; got %s", expected, got)
	}
}

func TestWarning(t *testing.T) {
	expected := sevWarning
	got := warning("Hello").Severity

	if expected != got {
		t.Errorf("expected %s; got %s", expected, got)
	}
}

func TestError(t *testing.T) {
	expected := sevError
	got := error("Hello").Severity

	if expected != got {
		t.Errorf("expected %s; got %s", expected, got)
	}
}

func TestFatal(t *testing.T) {
	expected := sevFatal
	got := fatal("Hello").Severity

	if expected != got {
		t.Errorf("expected %s; got %s", expected, got)
	}
}

func TestPrintLogItem(t *testing.T) {
	now, _ := time.Parse("2006-01-02T15:04:05.000000-07:00", "2019-06-28T08:13:00.591087-07:00")

	logItem := LogItem{
		Severity:  "debug",
		Timestamp: now,
		Message:   "Hello, World",
		Tags:      []Tag{"testing"},
		Caller: Caller{
			File: "logger_test.go",
			Line: 13,
			Func: "TestPrintLogItem",
		},
	}

	buffer := &bytes.Buffer{}
	dest = buffer

	printLogItem(logItem)

	expected := `{"severity":"debug","timestamp":"2019-06-28T08:13:00.591087-07:00","message":"Hello, World","tags":["testing"],"caller":{"file":"logger_test.go","line":13,"func":"TestPrintLogItem"}}`
	got := buffer.String()

	if expected != got {
		t.Errorf("expected %s; got %s", expected, got)
	}
}
