package database

import "cafe/models"

func migrate() error {
	err := DB.AutoMigrate(
		&models.User{},
	)
	return err
}
