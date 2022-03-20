package services

import "calendarapp/models"

type UserService interface {
	CreateUser(*models.User) error
	LoginUser(*string, *string) (*models.User)
	GetUser(*string) (*models.User)
	GetAllUsers() ([]*models.User)
	UpdateUser(*string, *models.User) (*models.User)
}