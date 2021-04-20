package database

import (
	"database/sql"

	"github.com/ben-eh/CodingOrganizer/entry"

	_ "github.com/go-sql-driver/mysql"
)

// func GetTasks() []task.Task {

// 	// Open up our database connection.
// 	// I've set up a database on my local machine using phpmyadmin.
// 	// The database is called testDb
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gojams")

// 	// if there is an error opening the connection, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// defer the close till after the main function has finished
// 	// executing
// 	defer db.Close()

// 	// Execute the query
// 	results, err := db.Query("SELECT id, name, completed FROM tasks")
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}

// 	var tasks []task.Task

// 	for results.Next() {
// 		var task task.Task
// 		// for each row, scan the result into our user composite object
// 		err = results.Scan(&task.ID, &task.Name, &task.Completed)
// 		if err != nil {
// 			panic(err.Error()) // proper error handling instead of panic in your app
// 		}
// 		tasks = append(tasks, task)
// 	}

// 	return tasks

// }

func SaveEntry(e entry.Entry) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/codingOrganizer")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Execute the query
	_, err2 := db.Query("INSERT INTO tasks SET name=?, url=?, codeblock=?, notes=?", e.Name, e.URL, e.Notes)
	if err2 != nil {
		panic(err2.Error()) // proper error handling instead of panic in your app
	}

}
