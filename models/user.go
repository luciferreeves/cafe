package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OpenID      string `gorm:"uniqueIndex;not null"`
	Username    string `gorm:"uniqueIndex;not null"`
	Email       string `gorm:"uniqueIndex;not null"`
	DisplayName string
	IsAdmin     bool `gorm:"default:false"`
}
