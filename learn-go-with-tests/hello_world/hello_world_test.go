package hello_world_test

import (
	"hello_world"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello world by fault.", func(t *testing.T) {
		want := "Hello, World!"
		got := hello_world.Hello("", "")
		
		assertCorrectMessage(t, want, got)
	})

	t.Run("Saying hello to people.", func(t *testing.T) {
		want := "Hello, Meowth!"
		got := hello_world.Hello("meowth", "")

		assertCorrectMessage(t, want, got)
	})

	t.Run("Saying hola to people.", func(t *testing.T) {
		want := "Hola, Ada!"
		got := hello_world.Hello("ada", "SPANISH")

		assertCorrectMessage(t, want, got)
	})

	t.Run("Saying bonjour to people.", func(t *testing.T) {
		want := "Bonjour, Wayne!"
		got := hello_world.Hello("wayne", "FRENCH")

		assertCorrectMessage(t, want, got)
	})
}


func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("Want %q; Got %q.", want, got)
	}
}
