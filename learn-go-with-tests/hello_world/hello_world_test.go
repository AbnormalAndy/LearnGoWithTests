package hello_world_test

import (
	"hello_world"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	want := "Hello, world!"
	got := hello_world.HelloWorld()
	
	if want != got {
		t.Errorf("Want %q; Got %q.", want, got)
	}
}


func TestHelloYou(t *testing.T) {
	want := "Hello, Meowth!"
	got := hello_world.HelloYou("meowth")

	if want != got {
		t. Errorf("Want %q; Got %q.", want, got)
	}
}
