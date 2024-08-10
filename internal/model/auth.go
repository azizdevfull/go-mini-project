package internal

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}
