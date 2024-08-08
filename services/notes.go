package services

import (
	"fmt"
	internal "go-tutorial/internal/model"

	"gorm.io/gorm"
)

type NoteService struct {
	db *gorm.DB
}

func (n *NoteService) InitNoteService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Note{})
}

func (n *NoteService) GetNoteService() []internal.Note {
	data := []internal.Note{
		{
			Id:     1,
			Title:  "Note 1",
			Status: true,
		},
	}
	return data
}

func (n *NoteService) CreateNoteService(title string, status bool) (*internal.Note, error) {
	data := &internal.Note{
		Title:  title,
		Status: status,
	}
	if err := n.db.Create(data).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}
	return data, nil
}
