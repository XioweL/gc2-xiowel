package handler_test

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"testing"
)

func setupTestDB() (*sql.DB, error) {
	// Membuka database SQLite untuk test
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to open test database: %v", err)
	}

	// Membuat tabel pengguna untuk test
	query := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);`

	if _, err := db.Exec(query); err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	return db, nil
}

func TestLoginUser(t *testing.T) {
	// Set up database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up test database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", "testuser", "password123")
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	// Cek jika query berhasil
	rows, err := db.Query("SELECT username FROM users WHERE username = ?", "testuser")
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}
	defer rows.Close()

	var username string
	for rows.Next() {
		if err := rows.Scan(&username); err != nil {
			t.Fatalf("failed to scan row: %v", err)
		}
		if username != "testuser" {
			t.Errorf("expected username 'testuser', got '%s'", username)
		}
	}
}
