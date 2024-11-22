package hello_world_test

import (
	"hello_world"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	want := "Hello, world!"
	got := hello_world.HelloWorld()
	
	if want != got {
		t.Errorf("Want %s; Got %s.", want, got)
	}
}
