package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to mohsen", func(t *testing.T) {

		got := Hello("Mohsen", English)
		want := EnglishHelloPrefix + "Mohsen"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty name provided", func(t *testing.T) {

		got := Hello("", English)
		want := EnglishHelloPrefix + "World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {

		got := Hello("", Spanish)
		want := SpanishHelloPrefix + "World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in franch", func(t *testing.T) {

		got := Hello("", Franch)
		want := FranchHelloPrefix + "World"

		assertCorrectMessage(t, got, want)
	})
}
