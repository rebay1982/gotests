package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got [%q] want [%q]", got, want)
		}
	}

	t.Run("Saying Hello to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello to World", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello to Elodie in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello to Elodie in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour Elodie"
		assertCorrectMessage(t, got, want)
	})

}
