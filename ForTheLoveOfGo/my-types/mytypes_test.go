package mytypes_test

import (
	"mytypes"
	"strings"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("Twice %d: Want %d; got %d.", input, want, got)
	}
}

func TestStringsBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("Want %q; got %q.", want, got)
	}
	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("%q: Want len %d; got %d.", sb.String(), wantLen, gotLen)
	}
}

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()
	if want != got {
		t.Errorf("Want %q; got %q.", want, got)
	}
}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()
	var su mytypes.StringUppercaser
	su.Contents.WriteString("Hello, Gophers!")
	want := "HELLO, GOPHERS!"
	got := su.ToUpper()
	if want != got {
		t.Errorf("Want %q; got %q.", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	var x int = 12
	want := 24
	mytypes.Double(&x)
	if want != x {
		t.Errorf("Want %d; got %d.", want, x)
	}
}
