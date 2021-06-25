package main

import "github.com/ben-eh/CodingOrganizer/pkg/models"

type templateData struct {
	Entry   *models.Entry
	Entries []*models.Entry
	Tag     *models.Tag
	Tags    []*models.Tag
}
