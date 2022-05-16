package event

import (
	"errors"
	"fmt"
	"net/http"
)

type Service interface {
	CreateNewEvent(input EventInput) (Event, error)
	UpdateEvent(eventID GetEventDetailInput, inputData EventInput) (Event, error)
	DeleteEvent(eventID int, userID int) (bool, int, string)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateNewEvent(input EventInput) (Event, error) {

	event := Event{}
	event.Title = input.Title
	event.Description = input.Description
	event.Perks = input.Perks
	event.StartDate = input.StartDate
	event.ClosingRegistration = input.ClosingRegistration
	event.LimitedTo = input.LimitedTo
	event.CompanyID = input.CompanyID
	event.UserID = input.User.ID
	event.Status = "on going"

	// TODO: fill CategoryID variabel

	newEvent, err := s.repository.CreateEvent(event)
	if err != nil {
		return newEvent, err
	}

	return newEvent, nil
}

func (s *service) UpdateEvent(eventID GetEventDetailInput, inputData EventInput) (Event, error) {

	event, err := s.repository.FindEventByID(eventID.ID)

	if err != nil {
		return event, err
	}

	if event.UserID != inputData.User.ID {
		return event, errors.New("not an owner of the event")
	}

	event.Title = inputData.Title
	event.Description = inputData.Description
	event.Perks = inputData.Perks
	event.StartDate = inputData.StartDate
	event.ClosingRegistration = inputData.ClosingRegistration
	event.LimitedTo = inputData.LimitedTo
	event.CompanyID = inputData.CompanyID
	event.UserID = inputData.User.ID
	event.Status = "on going"

	// TODO: fill CategoryID variabel

	updatedEvent, err := s.repository.UpdateEvent(event)
	if err != nil {
		return updatedEvent, err
	}

	return updatedEvent, nil
}

func (s *service) DeleteEvent(eventID int, userID int) (bool, int, string) {

	event, err := s.repository.FindEventByID(eventID)

	if err != nil {
		return false, http.StatusInternalServerError, err.Error()
	}

	if event.ID != int32(eventID) {
		message := fmt.Sprintf("Event with ID %d is not available", eventID)
		return false, http.StatusNotFound, errors.New(message).Error()
	}

	if event.UserID != userID {
		return false, http.StatusForbidden, errors.New("not an owner of the event").Error()
	}

	isDeleted, err := s.repository.DeleteEvent(int(event.ID))
	if err != nil {
		return false, http.StatusInternalServerError, err.Error()
	}

	return isDeleted, http.StatusOK, ""
}
