package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	want := 6.0
	got := calculator.Add(2.0, 4.0)
	if want != got {
		t.Errorf("Want: %f; got: %f.", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	want := 13.0
	got := calculator.Subtract(15.0, 2.0)
	if want != got {
		t.Errorf("Want: %f; Got: %f.", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	want := 20.0
	got := calculator.Multiply(4.0, 5.0)
	if want != got{
		t.Errorf("Want: %f; got: %f.", want, got)
	}
}