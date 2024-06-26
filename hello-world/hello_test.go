package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Fabio", "")
		want := "Hello, Fabio"

		assertCorrectMessage(got, want, t)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(got, want, t)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Holla, Elodie"

		assertCorrectMessage(got, want, t)
	})

}

func assertCorrectMessage(got string, want string, t testing.TB) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
