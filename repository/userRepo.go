package repository

import (
	"github.com/sadek-39/echo-sql/config"
	"github.com/sadek-39/echo-sql/models"
)

func GetUsersFromRepo() ([]models.User, error) {
	var DB = config.GetDBInstance()
	var users []models.User

	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
