package mysql

import (
	"database/sql"

	"github.com/ben-eh/CodingOrganizer/pkg/models"
)

type TagModel struct {
	DB *sql.DB
}

func (m *TagModel) GetAllTags() ([]*models.Tag, error) {
	stmt := `SELECT tag_id, name FROM tags`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tags := []*models.Tag{}

	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
