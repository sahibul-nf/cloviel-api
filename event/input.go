package event

import (
	"cloviel-api/user"
	"time"
)

type CompanyInput struct {
	Name             string `json:"name" binding:"required"`
	WebURL           string `json:"web_url"`
	ShortDescription string `json:"short_description"`
}

type CompanyLogoInput struct {
	CompanyID int `form:"company_id" binding:"required"`
}

type GetEventDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type EventInput struct {
	Title               string    `json:"title" binding:"required"`
	Description         string    `json:"description" binding:"required"`
	Perks               string    `json:"perks" binding:"required"`
	StartDate           time.Time `json:"start_date"`
	ClosingRegistration time.Time `json:"closing_registration"`
	LimitedTo           int32     `json:"limited_to"`
	CompanyID           int32     `json:"company_id"`
	User                user.User
}
