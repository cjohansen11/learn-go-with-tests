package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, exp string) {
		t.Helper()
		if got != exp {
			t.Errorf("got %q expected %q", got, exp)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Christian", "english")
		exp := "Hello Christian!"
		assertCorrectMessage(t, got, exp)
	})
	t.Run("say 'Hello World' when an empty string is provided", func(t *testing.T) {
		got := hello("", "english")
		exp := "Hello World!"
		assertCorrectMessage(t, got, exp)
	})
	t.Run("say hello in spanish", func(t *testing.T) {
		got := hello("Chi-Chi", "spanish")
		exp := "Hola Chi-Chi!"
		assertCorrectMessage(t, got, exp)
	})
	t.Run("say hello in french", func(t *testing.T) {
		got := hello("Christian", "french")
		exp := "Bonjour Christian!"
		assertCorrectMessage(t, got, exp)
	})
}
