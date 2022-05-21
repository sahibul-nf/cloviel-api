package event

import (
	"cloviel-api/company"
	"cloviel-api/presenter"
	"strings"
	"time"
)

type EventFormatter struct {
	ID                  int32                          `json:"id"`
	Title               string                         `json:"title"`
	Description         string                         `json:"description"`
	Perks               []string                       `json:"perks"`
	Status              string                         `json:"status"`
	MemberCount         int32                          `json:"member_count"`
	LimitedTo           int32                          `json:"limited_to"`
	Thumbnail           string                         `json:"thumbnail"`
	SignatureImage      string                         `json:"signature_image"`
	StartDate           time.Time                      `json:"start_date"`
	ClosingRegistration time.Time                      `json:"closing_registration"`
	CategoryID          int32                          `json:"category_id"`
	CompanyID           int32                          `json:"company_id"`
	UserID              int                            `json:"user_id"`
	User                UserEventFormatter             `json:"user"`
	Company             company.CompanyFormatter       `json:"company"`
	Presenters          []presenter.PresenterFormatter `json:"presenters"`
}

type UserEventFormatter struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func FormatDetailEvent(event Event) EventFormatter {

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

	user := event.User
	userEventFormatter := UserEventFormatter{}

	userEventFormatter.Name = user.Name
	userEventFormatter.Email = user.Email
	userEventFormatter.Avatar = user.AvatarFile

	formatter.User = userEventFormatter

	eventCompany := event.Company
	companyFormatter := company.CompanyFormatter{}

	companyFormatter.ID = eventCompany.ID
	companyFormatter.Name = eventCompany.Name
	companyFormatter.ShortDescription = eventCompany.ShortDescription
	companyFormatter.WebURL = eventCompany.WebURL
	companyFormatter.LogoURL = eventCompany.LogoURL

	formatter.Company = companyFormatter

	eventPresenters := event.Presenters
	presenterFormatters := []presenter.PresenterFormatter{}
	for _, v := range eventPresenters {
		presenterFormatters = append(presenterFormatters, presenter.FormatPresenter(v))
	}

	formatter.Presenters = presenterFormatters

	return formatter
}

func FormatEvents(events []Event) []EventFormatter {

	if len(events) == 0 {
		return []EventFormatter{}
	}

	var formatted []EventFormatter
	for _, event := range events {
		formatted = append(formatted, FormatDetailEvent(event))
	}

	return formatted
}
