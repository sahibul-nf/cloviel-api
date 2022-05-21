package handler

import "cloviel-api/presenter"

type presenterHandler struct {
	presenterService presenter.Service
}

func NewPresenterHandler(presenterService presenter.Service) *presenterHandler {
	return &presenterHandler{presenterService}
}

