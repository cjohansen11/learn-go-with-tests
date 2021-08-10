package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	exp := "Hello World!"

	if got != exp {
		t.Errorf("got %q expected %q", got, exp)
	}
}
