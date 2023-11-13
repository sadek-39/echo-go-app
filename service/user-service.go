package service

import (
	"github.com/sadek-39/echo-sql/models"
	"github.com/sadek-39/echo-sql/repository"
)

func GetUsers() ([]models.User, error) {
	users, _ := repository.GetUsersFromRepo()

	return users, nil
}
