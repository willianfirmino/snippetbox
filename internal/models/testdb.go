package models

import (
	"database/sql"
	"os"
	"testing"
	"time"
)

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "snippetbox_user:Passw0rd!0449@tcp(127.0.0.1:3306)/snippetbox_test_db?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	// Add a delay to give the database time to start up
	time.Sleep(5 * time.Second)

	script, err := os.ReadFile("./../../sql/schema.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_, err := db.Exec("DROP TABLE snippets; DROP TABLE users;")
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	})

	return db
}
