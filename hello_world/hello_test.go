package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHello(t *testing.T) {

	t.Run("test with param", func(t *testing.T) {
		var got string = Hello("Jonathan")
		var want string = "Hello, Jonathan"
		assertCorrectMessage(t, got, want)

	})
	t.Run("test without param", func(t *testing.T) {
		var got string = Hello("")
		var want string = "Hello, World"
		assertCorrectMessage(t, got, want)

	})
}
