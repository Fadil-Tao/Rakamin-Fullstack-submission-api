package models

import "gorm.io/gorm"

type Photos struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"url"`
	UserID   uint
	User     UserResponse `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type PhotosResponse struct {
	ID       uint   `json:"PhotoId"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"url"`
	UserID   uint   
}

func (PhotosResponse) TableName() string{
	return "photos"
}

func CreatePhoto(title, caption, photoUrl string, userid uint) (*Photos, error) {
	photo := &Photos{
		Title:    title,
		Caption:  caption,
		PhotoUrl: photoUrl,
		UserID:   userid,
	}

	err := DB.Create(&photo).Error
	if err != nil {
		return nil, err
	}
	return photo, nil
}
