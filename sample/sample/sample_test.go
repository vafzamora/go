package sample

import (
	"testing"
)

func TestResultString(t *testing.T) {
	if !(GetString("Teste") == "Teste is a good string.") {
		t.Fatal("Invalid string")
	}
}
