package services

import (
	"errors"
	internal "go-tutorial/internal/model"
	"go-tutorial/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	db.AutoMigrate(&internal.User{})
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) CheckUserExist(email *string) bool {
	var user internal.User
	if err := a.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return false
	}

	if user.Email != "" {
		return true
	}

	return false
}

func (a *AuthService) Login(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email cannot be empty")
	}

	if password == nil {
		return nil, errors.New("password cannot be empty")
	}
	var user internal.User
	if err := a.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	if user.Email == "" {
		return nil, errors.New("user not found")
	}
	if utils.CheckPasswordHash(*password, user.Password) == false {
		return nil, errors.New("password not match")
	}

	return &user, nil
}
func (a *AuthService) Register(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email cannot be empty")
	}

	if password == nil {
		return nil, errors.New("password cannot be empty")
	}

	if a.CheckUserExist(email) {
		return nil, errors.New("user already exist")
	}
	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return nil, err
	}

	var user internal.User
	user.Email = *email
	user.Password = hashedPassword
	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
