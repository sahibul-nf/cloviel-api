package presenter

import (
	"errors"
	"fmt"
	"net/http"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

type Service interface {
	CreateNewPresenter(input PresenterInput) (Presenter, error)
	UpdatePresenter(presenterID int, inputData PresenterInput) (Presenter, int, error)
}

func (s service) CreateNewPresenter(input PresenterInput) (Presenter, error) {

	presenter := Presenter{}
	presenter.Name = input.Name
	presenter.ShortDescription = input.ShortDescription
	presenter.EventID = input.EventID

	newPresenter, err := s.repository.CreatePresenter(presenter)
	if err != nil {
		return newPresenter, err
	}

	return newPresenter, nil
}

func (s service) UpdatePresenter(presenterID int, inputData PresenterInput) (Presenter, int, error) {

	presenter, err := s.repository.FindPresenterByID(presenterID)
	if err != nil {
		return presenter, http.StatusInternalServerError, err
	}

	if presenter.ID != presenterID {
		message := fmt.Sprintf("Presenter with ID %d is not available", presenterID)
		return Presenter{}, http.StatusNotFound, errors.New(message)
	}

	presenter.Name = inputData.Name
	presenter.ShortDescription = inputData.ShortDescription

	updatedPresenter, err := s.repository.UpdatePresenter(presenter)
	if err != nil {
		return updatedPresenter, http.StatusInternalServerError, err
	}

	return updatedPresenter, http.StatusOK, nil
}
