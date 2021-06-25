package mysql

import (
	"database/sql"

	"github.com/ben-eh/CodingOrganizer/pkg/models"
)

type EntryModel struct {
	DB *sql.DB
}

func (m *EntryModel) GetEntries() ([]*models.Entry, error) {

	stmt := `SELECT * FROM entries`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entries := []*models.Entry{}

	for rows.Next() {
		entry := &models.Entry{}
		err := rows.Scan(&entry.ID, &entry.Name, &entry.URL, &entry.CodeBlock, &entry.Notes)
		if err != nil {
			return nil, err
		}

		// entry.Tags = GetTagsForentry(entry)

		entries = append(entries, entry)
	}

	return entries, nil
}

// func dynamicSearch(r *http.Request) []Entry {

// 	db := database.DBConnection()
// 	defer db.Close()

// 	results, err2 := db.Query("SELECT * FROM entries WHERE ")
// 	if err2 != nil {
// 		panic(err2.Error())
// 	}

// 	vars := mux.Vars(r)
// 	log.Println(vars)
// 	log.Println("pause")

// }

// func SaveEntry(e Entry) int {
// 	stmt := `INSERT INTO entries SET (name, url, codeblock, notes) VALUES(?, ?, ?, ?)`
// 	db := database.DBConnection()
// 	defer db.Close()

// 	// Execute the query
// 	res, err2 := db.Exec("INSERT INTO entries SET name=?, url=?, codeblock=?, notes=?", e.Name, e.URL, e.CodeBlock, e.Notes)
// 	if err2 != nil {
// 		panic(err2.Error()) // proper error handling instead of panic in your app
// 	}

// 	lastID, err := res.LastInsertId()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return int(lastID)

// }

// func UpdateEntry(r *http.Request, e Entry) {
// 	log.Println("this is the first line in the UpdateEntry function in DB")
// 	db := database.DBConnection()
// 	defer db.Close()

// 	vars := mux.Vars(r)

// 	entryID := vars["entry_id"]

// 	log.Println("this should be the entry id: ", entryID)

// 	queryString := "UPDATE entries SET name=?, url=?, codeblock=?, notes=? WHERE entry_id=?"

// 	log.Println(queryString)

// 	// Execute the query
// 	result, err2 := db.Exec(queryString, e.Name, e.URL, e.CodeBlock, e.Notes, entryID)
// 	if err2 != nil {
// 		log.Println("there's an error executing the query in the updateEntry function")
// 		panic(err2.Error()) // proper error handling instead of panic in your app
// 	}
// 	log.Println(result)
// }

// func DeleteEntry(r *http.Request) {
// 	db := database.DBConnection()
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	entryID := vars["entry_id"]

// 	_, err := db.Query("DELETE FROM entries WHERE entry_id=?", entryID)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	log.Println("this is the entry ID in the delete Entry db function: ", entryID)
// }

// func FetchEntry(r *http.Request) Entry {
// 	db := database.DBConnection()
// 	defer db.Close()

// 	vars := mux.Vars(r)

// 	entryID := vars["entry_id"]
// 	// log.Println(reflect.TypeOf(entryID))
// 	log.Println("FetchEntry function in DB has this passed in http.Request for entry_id: ", entryID)

// 	results, err := db.Query("SELECT * FROM entries WHERE `entry_id`= '" + entryID + "'")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// log.Println(results)

// 	var entry Entry
// 	if results.Next() {
// 		err2 := results.Scan(&entry.ID, &entry.Name, &entry.URL, &entry.CodeBlock, &entry.Notes)
// 		if err2 != nil {
// 			panic(err2.Error())
// 		}
// 	}

// 	entry.Tags = GetTagsForentry(entry)

// 	return entry
// }

// func GetTagsForentry(e Entry) []Tag {

// 	db := database.DBConnection()
// 	defer db.Close()

// 	results, err2 := db.Query("SELECT tags.tag_id, tags.name from tags inner join entry_has_tag on (entry_has_tag.tag_id=tags.tag_id) where entry_has_tag.entry_id=?", e.ID)
// 	if err2 != nil {
// 		panic(err2.Error())
// 	}

// 	var tags []Tag

// 	for results.Next() {
// 		var tag Tag
// 		err := results.Scan(&tag.ID, &tag.Name)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		tags = append(tags, tag)
// 	}

// 	// log.Println(entries)
// 	// log.Println("pause")

// 	return tags
// }
