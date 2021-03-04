package main

import (
	"testing"
)

func TestOk(t *testing.T) {
	if true {
		t.Errorf("prueba test error")
	}
}
