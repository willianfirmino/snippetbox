package models

import (
	"testing"
)

func TestUserModel(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	t.Run("Authenticate", func(t *testing.T) {
		db := newTestDB(t)
		m := UserModel{DB: db}

		_, err := m.Authenticate("nonexistent@example.com", "password")
		if err != ErrInvalidCredentials {
			t.Errorf("want %v; got %v", ErrInvalidCredentials, err)
		}
	})
}
