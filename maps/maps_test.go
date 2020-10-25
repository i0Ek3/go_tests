package maps

import "testing"

func TestSearch(t *testing.T) {
	id := 1
	m := Maps{id: "map test"}

	t.Run("known id", func(t *testing.T) {
		got, _ := m.Search(id)
		want := "map test"

		assertStrings(t, got, want)
	})

	t.Run("unknown id", func(t *testing.T) {
		_, err := m.Search(-1)
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
