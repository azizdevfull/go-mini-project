package services

import (
	"errors"
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

func (n *NoteService) GetNoteService(status bool) ([]*internal.Note, error) {
	var data []*internal.Note

	if err := n.db.Where("status = ?", status).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (n *NoteService) CreateNoteService(title string, status bool) (*internal.Note, error) {
	data := &internal.Note{
		Title:  title,
		Status: status,
	}
	if data.Title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if err := n.db.Create(data).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}
	return data, nil
}
func (n *NoteService) UpdateNoteService(title string, status bool, id int) (*internal.Note, error) {

	var data *internal.Note

	if err := n.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	data.Title = title
	data.Status = status

	if err := n.db.Save(&data).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}
	return data, nil
}
func (n *NoteService) DeleteNoteService(id int64) error {

	var data *internal.Note

	if err := n.db.Where("id = ?", id).Delete(&data).Error; err != nil {
		return err
	}

	if err := n.db.Save(&data).Error; err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
