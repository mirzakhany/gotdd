package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to mohsen", func(t *testing.T) {

		got := Hello("Mohsen")
		want := HelloPrefix + "Mohsen"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty name provided", func(t *testing.T) {

		got := Hello("")
		want := HelloPrefix + "World"

		assertCorrectMessage(t, got, want)
	})
}
