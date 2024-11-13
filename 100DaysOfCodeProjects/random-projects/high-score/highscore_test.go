package highscore_test

import (
	"highscore"
	"testing"
)

func TestFindHighScore(t *testing.T) {
	t.Parallel()
	numbers := []int{78, 65, 89, 86, 55, 91, 64, 89}
	got := highscore.FindHighScore(numbers)
	want := 91
	if want != got {
		t.Errorf("Want %d; given %v; got %d.", want, numbers, got)
	}
}
