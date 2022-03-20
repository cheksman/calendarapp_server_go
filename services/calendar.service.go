package services

import "calendarapp/models"

type CalendarService interface {
	AddToCalendar(*models.Calendar) error
	GetUserCalendar(*string) ([]*models.Calendar, error)
	GetAllCalendars() ([]*models.Calendar, error)
	GetCalendarByInterval(*string) ([]*models.Calendar, error)
	EditCalendar(*string, *models.Calendar) (*models.Calendar, error)
}
