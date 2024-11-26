package integers_test


import (
	"fmt"
	"integers"
	"testing"
)


func TestAdd(t *testing.T) {
	t.Parallel()
	want := 3
	got := integers.Add(1, 2)
	if want != got {
		t.Errorf("Want %d; Got %d.", want, got)
	}
}


func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Print(sum)
	// Output: 6
}


