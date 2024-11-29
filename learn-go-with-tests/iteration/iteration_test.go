package iteration_test


import (
	"fmt"
	"iteration"
	"testing"
)


func TestRepeat(t *testing.T) {
	t.Parallel()
	want := "aaaaaaa"
	got := iteration.Repeat("a", 7)

	if want != got {
		t.Errorf("Want %s; Got %s.", want, got)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i <  b.N; i++ {
		iteration.Repeat("a", b.N)
	}
}


func ExampleRepeat() {
	repeated := iteration.Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}


