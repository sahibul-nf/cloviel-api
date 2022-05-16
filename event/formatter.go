package event

import (
	"strings"
	"time"
)

type EventFormatter struct {
	ID                  int32     `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Perks               []string  `json:"perks"`
	Status              string    `json:"status"`
	MemberCount         int32     `json:"member_count"`
	LimitedTo           int32     `json:"limited_to"`
	Thumbnail           string    `json:"thumbnail"`
	SignatureImage      string    `json:"signature_image"`
	StartDate           time.Time `json:"start_date"`
	ClosingRegistration time.Time `json:"closing_registration"`
	CategoryID          int32     `json:"category_id"`
	CompanyID           int32     `json:"company_id"`
	UserID              int       `json:"user_id"`
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

func FormatEvents(events []Event) []EventFormatter {

	if len(events) == 0 {
		return []EventFormatter{}
	}

	var formatted []EventFormatter
	for _, event := range events {
		formatted = append(formatted, FormatEvent(event))
	}

	return formatted
}