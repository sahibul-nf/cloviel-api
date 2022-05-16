package company

import "time"

type Company struct {
	ID               int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name             string    `gorm:"column:name;size:100"`
	ShortDescription string    `gorm:"column:short_description;size:255"`
	WebURL           string    `gorm:"column:web_url;size:255"`
	LogoURL          string    `gorm:"column:logo_url;size:255"`
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp"`
}
