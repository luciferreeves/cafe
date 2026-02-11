package repositories

import (
	"cafe/database"
	"cafe/models"
)

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func GetUserByOpenID(openID string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("open_id = ?", openID).First(&user).Error
	return &user, err
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}
