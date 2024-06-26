package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("João", "Portuguese")
		want := "Olá, João"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
