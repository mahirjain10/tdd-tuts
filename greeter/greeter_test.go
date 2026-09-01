package greeter

import "testing"

func TestGreet(t *testing.T) {
	t.Run("If name and language is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})
	t.Run("If name is given and language is not given", func(t *testing.T) {
		got := Hello("Mahir", "")
		want := "Hello, Mahir"
		assertMessage(t, got, want)
	})
	t.Run("If the langauge is English", func(t *testing.T) {
		got := Hello("Mahir", "English")
		want := "Hello, Mahir"
		assertMessage(t, got, want)
	})
	t.Run("If the langauge is Spanish", func(t *testing.T) {
		got := Hello("Mahir", spanish)
		want := "Hola, Mahir"
		assertMessage(t, got, want)
	})
	t.Run("If the langauge is French", func(t *testing.T) {
		got := Hello("Mahir", french)
		want := "Bonjour, Mahir"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s and want %s", got, want)
	}
}
