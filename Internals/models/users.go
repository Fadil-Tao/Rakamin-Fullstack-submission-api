package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"ID"`
	Username  string `gorm:"not null" json:"username"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos    []PhotosResponse
}

type UserResponse struct{
	ID        uint   `json:"ID"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func (UserResponse) TableName() string{
	return "users"
}

func UserAvailable(email string, username string) bool {
	var user User
	DB.Where("email = ?", email).Or("username = ?", username).First(&user)
	return user.ID == 0
}

func CreateUser(email, username, password string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

