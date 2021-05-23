package entryHasTag

import (
	"log"

	"github.com/ben-eh/CodingOrganizer/database"
)

func CreateEntryHasTag(tags []string, lastID int) {
	db := database.DBConnection()
	defer db.Close()

	// get most recenty entry_id
	// lastID, _ := db.Exec("SELECT LAST_INSERT_ID();")
	// log.Println(reflect.TypeOf(lastID))
	log.Println(lastID)
	log.Println(tags)
	log.Println("pause")

	// Execute the query
	for _, tag := range tags {
		_, err2 := db.Query("INSERT INTO entry_has_tag SET entry_id=?, tag_id=?", lastID, tag)
		if err2 != nil {
			panic(err2.Error()) // proper error handling instead of panic in your app
		}
	}

}
