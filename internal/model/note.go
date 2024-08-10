package model

type Note struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (Note) TableName() string {
	return "notes"
}
