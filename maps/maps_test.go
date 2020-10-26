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

func TestAdd(t *testing.T) {
	t.Run("new value", func(t *testing.T) {
		m := Maps{}
		id := 2
		value := "test too"
		err := m.Add(id, value)

		assertError(t, err, nil)
		assertValue(t, m, id, value)
	})

	t.Run("existing value", func(t *testing.T) {
		m := Maps{}
		id := 1
		value := "map test"
		err := m.Add(id, value)

		assertError(t, err, ErrExist)
		assertValue(t, m, id, value)
	})
}

func assertValue(t *testing.T, m Maps, id int, value string) {
	t.Helper()

	got, err := m.Search(id)
	if err != nil {
		t.Fatal("should find added value:", err)
	}

	if value != got {
		t.Errorf("got '%s' want '%s'", got, value)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("new value", func(t *testing.T) {
		id := 3
		value := "test too"
		m := Maps{}

		err := m.Update(id, value)

		assertError(t, err, ErrNotExist)
	})

	t.Run("existing value", func(t *testing.T) {
		id := 3
		value := "test too"
		m := Maps{id: value}
		newValue := "new value"

		err := m.Update(id, newValue)

		assertError(t, err, nil)
		assertValue(t, m, id, newValue)
	})
}

func TestDelete(t *testing.T) {
	t.Run("not existing value", func(t *testing.T) {
		id := 4
		m := Maps{}

		err := m.Delete(id)

		assertError(t, err, ErrNotExist)
	})

	t.Run("existing value", func(t *testing.T) {
		id := 3
		m := Maps{}

		err := m.Delete(id)

		assertError(t, err, nil)
	})
}
