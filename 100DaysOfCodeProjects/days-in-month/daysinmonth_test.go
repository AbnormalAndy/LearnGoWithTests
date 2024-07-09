package daysinmonth_test

import (
	"daysinmonth"
	"testing"
)

func TestDaysInMonth(t *testing.T) {
	t.Parallel()
	want := 28
	got := daysinmonth.DaysInMonth(2022, 2)
	if got != want {
		t.Errorf("Got %d; want %d.", got, want)
	}
}

