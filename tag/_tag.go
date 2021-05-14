package tag

import (
	"github.com/ben-eh/CodingOrganizer/database"
	"github.com/ben-eh/CodingOrganizer/entry"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetTagsForentry(e entry.Entry) []Tag {

	db := database.DBConnection()
	defer db.Close()

	results, err2 := db.Query("SELECT tags.* from tags inner join entry_has_tag on (entry_has_tag.tag_id=tags.tag_id) where entry_has_tag.entry_id=?", e.ID)
	if err2 != nil {
		panic(err2.Error())
	}

	var tags []Tag

	for results.Next() {
		var tag Tag
		err := results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		tags = append(tags, tag)
	}

	// log.Println(entries)
	// log.Println("pause")

	return tags
}
