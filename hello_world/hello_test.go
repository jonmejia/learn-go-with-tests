package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHello(t *testing.T) {

	t.Run("test with English param", func(t *testing.T) {
		var got string = Hello("Jonathan", "English")
		var want string = "Hello, Jonathan"
		assertCorrectMessage(t, got, want)

	})
	t.Run("test with Spanish param", func(t *testing.T) {
		var got string = Hello("Jonathan", "Spanish")
		var want string = "Hola, Jonathan"
		assertCorrectMessage(t, got, want)

	})
	t.Run("test with French param", func(t *testing.T) {
		var got string = Hello("Jonathan", "French")
		var want string = "Bonjour, Jonathan"
		assertCorrectMessage(t, got, want)

	})

	t.Run("test without param", func(t *testing.T) {
		var got string = Hello("", "English")
		var want string = "Hello, World"
		assertCorrectMessage(t, got, want)

	})
}
