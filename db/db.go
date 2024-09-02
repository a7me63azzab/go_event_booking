package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 email TEXT NOT NULL UNIQUE,
	 password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

}

// package db

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DB *sql.DB

// func InitDB() {
// var err error
// 	DB, err = sql.Open("sqlite3", "./event_booking.db")

// 	if err != nil {
// 		// panic("Could not connect to database")
// 		log.Fatalf("Error opening database: %v", err)
// 	}

// 	DB.SetMaxOpenConns(10)
// 	DB.SetMaxIdleConns(5)

// 	// Verify the connection is valid
// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatalf("Error connecting to the database: %v", err)
// 	}

// 	createTables()
// }

// func createTables() {

// 	if DB == nil {
// 		log.Fatal("DB is not initialized")
// 	}

// 	createEventsTable := `
// 		CREATE TABLE IF NOT EXISTS events (
// 		 id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		 name TEXT NOT NULL,
//          description TEXT NOT NULL,
// 		 location TEXT NOT NULL,
// 		 dateTime DATETIME NOT NULL,
// 		 user_id INTEGER
// 		)
// 	`

// 	_, err := DB.Exec(createEventsTable)

// 	if err != nil {
// 		panic("Could not create events table.")
// 	}
// }
