package main

import "testing"
import "bytes"

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ellery")

	got := buffer.String()
	want := "Hello Ellery"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}