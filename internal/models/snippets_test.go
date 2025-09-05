package models

import (
	"testing"
)

func TestSnippetModel(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	t.Run("Insert", func(t *testing.T) {
		db := newTestDB(t)
		m := SnippetModel{DB: db}

		_, err := m.Insert("Test Snippet", "This is a test snippet", 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Get", func(t *testing.T) {
		db := newTestDB(t)
		m := SnippetModel{DB: db}

		id, err := m.Insert("Test Snippet", "This is a test snippet", 1)
		if err != nil {
			t.Fatal(err)
		}

		s, err := m.Get(id)
		if err != nil {
			t.Fatal(err)
		}

		if s.Title != "Test Snippet" {
			t.Errorf("want %q; got %q", "Test Snippet", s.Title)
		}
	})

	t.Run("Latest", func(t *testing.T) {
		db := newTestDB(t)
		m := SnippetModel{DB: db}

		_, err := m.Insert("Test Snippet 1", "This is a test snippet", 1)
		if err != nil {
			t.Fatal(err)
		}
		_, err = m.Insert("Test Snippet 2", "This is a test snippet", 1)
		if err != nil {
			t.Fatal(err)
		}

		snippets, err := m.Latest()
		if err != nil {
			t.Fatal(err)
		}

		if len(snippets) != 2 {
			t.Errorf("want %d; got %d", 2, len(snippets))
		}
	})
}
