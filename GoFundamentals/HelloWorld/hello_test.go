package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people.", func(t *testing.T) {
		got := Hello("Meow", "")
		want := "Hello, Meow."
		assertCorrectMessage(t, got, want)
	})
	t.Run("Empty string defaults to 'world'.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world."
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Meow", "French")
		want := "Bonjour, Meow."
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Meow", "Russian")
		want := "Привет, Meow."
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie."
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q; want %q.", got, want)
	}
}
