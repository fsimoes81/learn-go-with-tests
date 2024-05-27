package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Fabio")
	want := "Hello, Fabio"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
