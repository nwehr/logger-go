package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	defer func() {
		now = func() time.Time {
			return time.Now()
		}

		dest = os.Stdout
	}()

	now = func() time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05.000000-07:00", "2019-06-28T08:13:00.591087-07:00")
		return t
	}

	buffer := &bytes.Buffer{}
	dest = buffer

	Debug("Hello, World", "testing")

	expected := `{"severity":"debug","timestamp":"2019-06-28T08:13:00.591087-07:00","message":"Hello, World","tags":["testing"],`
	got := buffer.String()

	if !strings.Contains(got, expected) {
		t.Errorf("expected %s; got %s", expected, got)
	}
}

func TestWarning(t *testing.T) {
	defer func() {
		now = func() time.Time {
			return time.Now()
		}

		dest = os.Stdout
	}()

	now = func() time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05.000000-07:00", "2019-06-28T08:13:00.591087-07:00")
		return t
	}

	buffer := &bytes.Buffer{}
	dest = buffer

	Warning("Hello, World", "testing")

	expected := `{"severity":"warning","timestamp":"2019-06-28T08:13:00.591087-07:00","message":"Hello, World","tags":["testing"],`
	got := buffer.String()

	if !strings.Contains(got, expected) {
		t.Errorf("expected %s; got %s", expected, got)
	}
}

func TestError(t *testing.T) {
	defer func() {
		now = func() time.Time {
			return time.Now()
		}

		dest = os.Stdout
	}()

	now = func() time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05.000000-07:00", "2019-06-28T08:13:00.591087-07:00")
		return t
	}

	buffer := &bytes.Buffer{}
	dest = buffer

	Error("Hello, World", "testing")

	expected := `{"severity":"error","timestamp":"2019-06-28T08:13:00.591087-07:00","message":"Hello, World","tags":["testing"],`
	got := buffer.String()

	if !strings.Contains(got, expected) {
		t.Errorf("expected %s; got %s", expected, got)
	}
}

func TestFatal(t *testing.T) {
	defer func() {
		now = func() time.Time {
			return time.Now()
		}

		exit = func(exitCode int) {
			os.Exit(exitCode)
		}

		dest = os.Stdout
	}()

	exit = func(exitCode int) {
		_ = exitCode
	}

	now = func() time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05.000000-07:00", "2019-06-28T08:13:00.591087-07:00")
		return t
	}

	buffer := &bytes.Buffer{}
	dest = buffer

	Fatal("Hello, World", "testing")

	expected := `{"severity":"fatal","timestamp":"2019-06-28T08:13:00.591087-07:00","message":"Hello, World","tags":["testing"],`
	got := buffer.String()

	if !strings.Contains(got, expected) {
		t.Errorf("expected %s; got %s", expected, got)
	}
}
