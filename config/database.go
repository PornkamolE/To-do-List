package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite" //ใช้ sqlite
)

func Database() *sql.DB {
	db, err := sql.Open("sqlite", "gotodo.db") 
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SQLite Database Connection Successful")
	}

	// สร้างตารางถ้ายังไม่มี
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
