package internal

type Note struct {
	Id     int    `gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
