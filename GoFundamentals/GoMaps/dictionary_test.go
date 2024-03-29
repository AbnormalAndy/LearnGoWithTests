package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test."}

	t.Run("Known Word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test."

		assertStrings(t, got, want)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New word.", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is just a test."

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word.", func(t *testing.T) {
		word := "test"
		definition := "This is just a test."
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "New test.")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word.", func(t *testing.T) {
		word := "test"
		definition := "This is just a test."
		dictionary := Dictionary{word: definition}
		newDefinition := "New definition."

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("New word.", func(t *testing.T) {
		word := "test"
		definition := "This is just a test."
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "Test definition."}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted.", word)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q; want %q.", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got error %q; want %q.", got, want)
	}
}
