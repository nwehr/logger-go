package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	logItem := debug("Hello, World")

	if logItem.Severity != DEBUG {
		t.Errorf("expected severity %s; got %s", DEBUG, logItem.Severity)
	}
}

func TestWarning(t *testing.T) {
	logItem := warning("Hello, World")

	if logItem.Severity != WARNING {
		t.Errorf("expected severity %s; got %s", WARNING, logItem.Severity)
	}
}

func TestError(t *testing.T) {
	logItem := error("Hello, World")

	if logItem.Severity != ERROR {
		t.Errorf("expected severity %s; got %s", ERROR, logItem.Severity)
	}
}

func TestFatal(t *testing.T) {
	logItem := fatal("Hello, World")

	if logItem.Severity != FATAL {
		t.Errorf("expected severity %s; got %s", FATAL, logItem.Severity)
	}
}
