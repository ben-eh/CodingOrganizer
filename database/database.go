package database

import (
	"database/sql"

	"github.com/ben-eh/CodingOrganizer/entry"

	_ "github.com/go-sql-driver/mysql"
)

func GetEntries() []entry.Entry {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/codingOrganizer")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err2 := db.Query("SELECT * FROM entries")
	if err2 != nil {
		panic(err2.Error())
	}

	var entries []entry.Entry

	for results.Next() {
		var entry entry.Entry
		err = results.Scan(&entry.ID, &entry.Name, &entry.URL, &entry.Notes)
		if err != nil {
			panic(err.Error())
		}
		entries = append(entries, entry)
	}

	return entries

}

func SaveEntry(e entry.Entry) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/codingOrganizer")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Execute the query
	_, err2 := db.Query("INSERT INTO entries SET name=?, url=?, codeblock=?, notes=?", e.Name, e.URL, e.CodeBlock, e.Notes)
	if err2 != nil {
		panic(err2.Error()) // proper error handling instead of panic in your app
	}
}
