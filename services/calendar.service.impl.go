package services

import (
	"calendarapp/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CalendarServiceImpl struct {
	calendarCollection *mongo.Collection
	ctx                context.Context
}

func CalendarServiceImplementation(calendarCollection *mongo.Collection, ctx context.Context) CalendarService {
	return &CalendarServiceImpl{
		calendarCollection: calendarCollection,
		ctx: ctx,
	}
}

func (c *CalendarServiceImpl) AddToCalendar(calendar *models.Calendar) (*models.Calendar, error) {
	return nil, nil
}

func (c *CalendarServiceImpl) GetUserCalendar(userId *string) ([]*models.Calendar, error) {
	return nil, nil
}

func (c *CalendarServiceImpl) GetAllCalendars() ([]*models.Calendar, error) {
	return nil, nil
}

func (c *CalendarServiceImpl) GetCalendarByInterval(interval *string) ([]*models.Calendar, error) {
	return nil, nil
}

func (c *CalendarServiceImpl) EditCalendar(userId *string, calendar *models.Calendar) (*models.Calendar, error) {
	return nil, nil
}