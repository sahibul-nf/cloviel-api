package event

import "time"

type Event struct {
	ID                  int32     `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID              int32     `gorm:"column:user_id;type:int"`
	CategoryID          int32     `gorm:"column:category_id;type:int32"`
	CompanyID           int32     `gorm:"column:company_id;type:int"`
	Title               string    `gorm:"column:title;size:255"`
	Description         string    `gorm:"column:description;size:255"`
	Perks               string    `gorm:"column:perks;size:255"`
	Status              string    `gorm:"column:status;size:255"`
	MemberCount         int32     `gorm:"column:member_count;type:int"`
	LimitedTo           int32     `gorm:"column:limited_to;type:int"`
	Thumbnail           string    `gorm:"column:thumbnail;size:255"`
	SignatureImage      string    `gorm:"column:signature_image;size:255"`
	StartDate           time.Time `gorm:"column:start_date;type:timestamp"`
	ClosingRegistration time.Time `gorm:"column:closing_registration;type:timestamp"`
	CreatedAt           time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:timestamp"`
}

type Company struct {
	ID               int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name             string    `gorm:"column:name;size:100"`
	ShortDescription string    `gorm:"column:short_description;size:255"`
	WebURL           string    `gorm:"column:web_url;size:255"`
	LogoURL          string    `gorm:"column:logo_url;size:255"`
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp"`
}
