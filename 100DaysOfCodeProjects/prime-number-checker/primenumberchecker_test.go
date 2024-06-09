package primenumberchecker_test

import (
	"primenumberchecker"
	"testing"
)

func TestPrimeNumberChecker(t *testing.T) {
	t.Parallel()
	want := "It is a prime number."
	got := primenumberchecker.PrimeNumberChecker(73)
	if want != got {
		t.Errorf("Want '%s'; got '%s'.", want, got)
	}
}
