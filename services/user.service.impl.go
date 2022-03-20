package services

import (
	"calendarapp/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func UserServiceImplementation(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) LoginUser(email *string, password *string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) GetUser(userId *string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(userId *string, user *models.User) (*models.User, error) {
	return nil, nil
}
