package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello, world"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	t.Run("greetings with name", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("generic greetings for empty function", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
