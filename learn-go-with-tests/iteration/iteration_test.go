package iteration_test


import (
	"iteration"
	"testing"
)


func TestRepeat(t *testing.T) {
	t.Parallel()
	want := "aaaaa"
	got := iteration.Repeat("a")

	if want != got {
		t.Errorf("Want %s; Got %s.", want, got)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i <  b.N; i++ {
		iteration.Repeat("a")
	}
}
