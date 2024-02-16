package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Meow")
	want := "Hello, Meow."

	if got != want {
		t.Errorf("Got %q; want %q.", got, want)
	}
}
