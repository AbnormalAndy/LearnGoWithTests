package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	input := calculator.Add(2, 4)
	got := input
	want := 6
	if got != want {
		t.Errorf("Input: %d; Got: %d; Want: %d.", input, got, want)
	}
}
