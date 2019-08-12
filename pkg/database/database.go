package database

import (
	"database/sql"
	"fmt"
	"strconv"
)

// TODO: create a config file for the database.

func openDB() (database *sql.DB) {
	database, _ = sql.Open("sqlite3", "./gft.db")
	defer database.Close()

	// Create db if not there already
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS files (id INTEGER PRIMARY KEY,from_add TEXT, local_file TEXT, tx_address TEXT, tx_data TEXT, file_hash TEXT, url TEXT, rsa_key TEXT)")
	statement.Exec()
	defer statement.Close()
	fmt.Println("Opened DB successfuly")
	return
}

func RecordFile(tx_address string, tx_data string) {
	db, _ := sql.Open("sqlite3", "./gft.db")
	defer db.Close()
	statement, _ := db.Prepare("INSERT INTO files (tx_address, tx_data) VALUES (?, ?)")
	statement.Exec(tx_address, tx_data)
	defer statement.Close()
}

func ListAllFiles() {
	db, _ := sql.Open("sqlite3", "./gft.db")
	defer db.Close()
	rows, _ := db.Query("SELECT id, tx_data FROM files")
	var id int
	var txData string
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &txData)
		fmt.Println(strconv.Itoa(id) + ": " + txData)
	}
}

func ListOneFile(id int) {
	fmt.Println("Test")
}
