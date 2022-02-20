package event

import "time"

type Event struct {
	ID                  int
	UserID              int
	CategoryID          int
	CompanyID           int
	Title               string
	Description         string
	Perks               string
	Status              string
	MemberCount         int
	LimitedTo           int
	Thumbnail           string
	SignatureImage      string
	StartDate           time.Time
	ClosingRegistration time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Company struct {
	ID               int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Name             string    `gorm:"column:name;size:100" json:"name"`
	ShortDescription string    `gorm:"column:short_description;size:255" json:"short_description"`
	WebURL           string    `gorm:"column:web_url;size:255" json:"web_url"`
	LogoURL          string    `gorm:"column:logo_url;size:255" json:"logo_url"`
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}
