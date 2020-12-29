package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func (t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	// t.Run("say hello to peopel", func (t *testing.T) {
	// 	got := Hello("Ellery")
	// 	want := "Hello Ellery!"

	// 	assertMessage(t, got, want)
	// })

	// t.Run("say hello world when an empty string is supplied", func (t *testing.T) {
	// 	got := Hello("")
	// 	want := "Hello World!"

	// 	assertMessage(t, got, want)
	// })

	t.Run("in chinese", func (t *testing.T) {
		got := Hello("Ellery", "Chinese")
		want := "你好 Ellery!"

		assertMessage(t, got, want)
	})
}