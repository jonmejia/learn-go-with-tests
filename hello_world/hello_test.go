package main

import "testing"

func TestHello(t *testing.T) {
	var got string = Hello("Jonathan")
	var want string = "Hello, Jonathan"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
