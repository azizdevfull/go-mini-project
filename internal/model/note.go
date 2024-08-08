package internal

type Note struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
