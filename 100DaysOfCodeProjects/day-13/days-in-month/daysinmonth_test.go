package daysinmonth_test

import (
	"daysinmonth"
	"testing"
)

func TestIsLeap(t *testing.T) {
	t.Parallel()
	want := true
	got := daysinmonth.IsLeap(2000)
	if got != want {
		t.Errorf("Got %v; want %v.", got, want)
	}
}

func TestDaysInMonth(t *testing.T) {
	t.Parallel()
	want := 28
	got := daysinmonth.DaysInMonth(2022, 2)
	if got != want {
		t.Errorf("Got %d; want %d.", got, want)
	}
}

