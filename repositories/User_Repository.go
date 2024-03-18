package repositories

import "lets-go/models"

type UserRepository interface {
	GetUsers() ([]models.User, error)
}
