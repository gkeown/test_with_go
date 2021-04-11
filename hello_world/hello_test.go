package main

import "testing"

func TestGreeting(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello in English", func(t *testing.T) {
		got := Greeting("Gerard", "English")
		want := "Hello, Gerard!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Greeting("Paulo", "Spanish")
		want := "Hola, Paulo!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Greeting("Pierre", "French")
		want := "Bonjour, Pierre!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Greeting("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}
