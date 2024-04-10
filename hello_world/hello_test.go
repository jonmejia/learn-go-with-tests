package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("test with param", func(t *testing.T) {
		var got string = Hello("Jonathan")
		var want string = "Hello, Jonathan"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
	t.Run("test without param", func(t *testing.T) {
		var got string = Hello("")
		var want string = "Hello, World"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
}
