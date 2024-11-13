package helloworld_test

import (
	"helloworld"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	want := "Hello, Meow."
	got := helloworld.HelloWorld("Meow")
	if want != got {
		t.Errorf("Got %q; want %q.", want, got)
	}
}
