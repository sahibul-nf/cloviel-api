package presenter

import "cloviel-api/user"

type PresenterInput struct {
	EventID          int    `json:"event_id" binding:"required"`
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	User             user.User
}

type PresenterDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
