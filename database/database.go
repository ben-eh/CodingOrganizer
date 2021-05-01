package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/ben-eh/CodingOrganizer/entry"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/codingOrganizer")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetEntries() []entry.Entry {

	db := DBConnection()
	defer db.Close()

	results, err2 := db.Query("SELECT * FROM entries")
	if err2 != nil {
		panic(err2.Error())
	}

	var entries []entry.Entry

	for results.Next() {
		var entry entry.Entry
		err := results.Scan(&entry.ID, &entry.Name, &entry.URL, &entry.CodeBlock, &entry.Notes)
		if err != nil {
			panic(err.Error())
		}
		entries = append(entries, entry)
	}

	// log.Println(entries)
	// log.Println("pause")

	return entries
}

func SaveEntry(e entry.Entry) {
	db := DBConnection()
	defer db.Close()

	// Execute the query
	_, err2 := db.Query("INSERT INTO entries SET name=?, url=?, codeblock=?, notes=?", e.Name, e.URL, e.CodeBlock, e.Notes)
	if err2 != nil {
		panic(err2.Error()) // proper error handling instead of panic in your app
	}
}

func UpdateEntry(r *http.Request, e entry.Entry) {
	log.Println("this is the first line in the UpdateEntry function in DB")
	db := DBConnection()
	defer db.Close()

	vars := mux.Vars(r)

	entryID := vars["entry_id"]

	log.Println("this should be the entry id: ", entryID)

	queryString := fmt.Sprintf("UPDATE entries SET name=?, url=?, codeblock=?, notes=? WHERE entry_id='%v'", entryID)

	log.Println(queryString)

	// Execute the query
	result, err2 := db.Exec(queryString, e.Name, e.URL, e.CodeBlock, e.Notes, entryID)
	if err2 != nil {
		log.Println("there's an error executing the query in the updateEntry function")
		panic(err2.Error()) // proper error handling instead of panic in your app
	}
	log.Println(result)
}

func FetchEntry(r *http.Request) entry.Entry {
	db := DBConnection()
	defer db.Close()

	vars := mux.Vars(r)

	entryID := vars["entry_id"]
	// log.Println(reflect.TypeOf(entryID))
	log.Println("FetchEntry function in DB has this passed in http.Request for entry_id: ", entryID)

	results, err := db.Query("SELECT * FROM entries WHERE `entry_id`= '" + entryID + "'")
	if err != nil {
		panic(err.Error())
	}
	// log.Println(results)

	var entry entry.Entry
	if results.Next() {
		err2 := results.Scan(&entry.ID, &entry.Name, &entry.URL, &entry.CodeBlock, &entry.Notes)
		if err2 != nil {
			panic(err2.Error())
		}
	}
	// log.Println(entry)
	return entry
}
