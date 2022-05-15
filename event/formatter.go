package event

import (
	"strings"
	"time"
)

type CompanyFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	WebURL           string `json:"web_url"`
	LogoURL          string `json:"logo_url"`
	ShortDescription string `json:"short_description"`
}

func FormatCompany(company Company) CompanyFormatter {

	formatter := CompanyFormatter{}
	formatter.ID = company.ID
	formatter.Name = company.Name
	formatter.WebURL = company.WebURL
	formatter.LogoURL = company.LogoURL
	formatter.ShortDescription = company.ShortDescription

	return formatter
}

type EventFormatter struct {
	ID                  int32            `json:"id"`
	Title               string           `json:"title"`
	Description         string           `json:"description"`
	Perks               []string         `json:"perks"`
	Status              string           `json:"status"`
	MemberCount         int32            `json:"member_count"`
	LimitedTo           int32            `json:"limited_to"`
	Thumbnail           string           `json:"thumbnail"`
	SignatureImage      string           `json:"signature_image"`
	StartDate           time.Time        `json:"start_date"`
	ClosingRegistration time.Time        `json:"closing_registration"`
	CategoryID          int32            `json:"category_id"`
	CompanyID           int32            `json:"company_id"`
	UserID              int            `json:"user_id"`
}

type UserEventFormatter struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func FormatEvent(event Event) EventFormatter {

	formatter := EventFormatter{}
	formatter.ID = event.ID
	formatter.CategoryID = event.CategoryID
	formatter.Title = event.Title
	formatter.Description = event.Description
	formatter.Thumbnail = event.Thumbnail
	formatter.StartDate = event.StartDate
	formatter.ClosingRegistration = event.ClosingRegistration
	formatter.MemberCount = event.MemberCount
	formatter.LimitedTo = event.LimitedTo
	formatter.CompanyID = event.CompanyID
	formatter.UserID = event.UserID
	formatter.Status = event.Status

	var perks []string
	for _, perk := range strings.Split(event.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	formatter.Perks = perks

	return formatter
}
