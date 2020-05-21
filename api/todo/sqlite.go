// Package todo uses sqlite3 driver to GET and PUSH to the database from the user
package todo

import (
	"database/sql"
	"log"
)

// SQLite holds the DB connection
type SQLite struct {
	DB *sql.DB
}

// Set stores items in the database
func (conn *SQLite) Set(item Item) bool {
	stmt, err := conn.DB.Prepare("INSERT INTO todo_items (text, complete) VALUES (?, ?)")
	CheckError(err)
	stmt.Exec(item.Text, item.Complete)
	return true
}

// Get returns all todo items from the todo_items table
func (conn *SQLite) Get() []Item {
	items := []Item{}
	rows, err := conn.DB.Query("SELECT * FROM todo_items")
	CheckError(err)
	defer rows.Close()
	var id int
	var text string
	var complete int
	for rows.Next() {
		rows.Scan(&id, &text, &complete)
		items = append(items, Item{id, text, complete})
	}
	return items
}

// Update returns a bool and updates an item on todo_items
func (conn *SQLite) Update(item Item) bool {
	stmt, err := conn.DB.Prepare("UPDATE todo_items SET complete=? WHERE id=?")
	CheckError(err)
	stmt.Exec(item.Complete, item.ID)
	return true
}

// GetDBConnection creates initial todo_items table and returns connection
func GetDBConnection(conn *sql.DB) *SQLite {
	stmt, err := conn.Prepare(`
		CREATE TABLE IF NOT EXISTS todo_items (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			text TEXT,
			complete INTEGER
		);
	`)
	CheckError(err)
	
	stmt.Exec()
	return &SQLite{
		DB: conn,
	}
}

// CheckError takes an error and checks if its nil
// This probably belongs in some other package
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}