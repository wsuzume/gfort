package gfort

import (
	"errors"
	//"fmt"
	"testing"
	//"os"

	//"github.com/pkg/errors"
)

func success() (string, error) {
	return "test", nil
}

func failure() (string, error) {
	return "", errors.New("simulated failure")
}

func TestIgnore(t *testing.T) {
	ActivateWarning()
	Ignore(nil)
}

func TestIgnoreSuppressed(t *testing.T) {
	SuppressWarning()
	Ignore(nil)
}

func TestPanicWithSuccess(t *testing.T) {
	ActivateWarning()
	PanicF(success())
}

func TestPanicWithFailure(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("The code must cause panic but did not.")
		}
	}()

	ActivateWarning()
	PanicF(failure())
}

func TestFilter(t *testing.T) {
	ActivateWarning()

	s := Filter(IgnoreF, success)
	println(s.(string))
}
