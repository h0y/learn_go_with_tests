package maps

import "testing"

func TestSearch(t *testing.T) {
	// dictionary := map[string]string {"test": "this is a test"}
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func (t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func (t *testing.T) {
		_, err := dictionary.Search("test2")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	word := "test2"
	definition := "this is just test2"
	dictionary.Add(word, definition)
	
	assertDefinition(t, dictionary, word, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search("test2")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}