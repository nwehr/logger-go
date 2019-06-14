package logger

import (
	"testing"
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
