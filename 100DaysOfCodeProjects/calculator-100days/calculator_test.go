package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	want := 6.0
	got := calculator.Add(2.0, 4.0)
	if got != want {
		t.Errorf("Got: %f; Want: %f.", got, want)
	}
}
