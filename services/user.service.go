package services

import "calendarapp/models"

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
	LoginUser(*string, *string) (*models.User, error)
	GetUser(*string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(*string, *models.User) (*models.User, error)
}
