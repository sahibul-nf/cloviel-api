package presenter

import "time"

type Presenter struct {
	ID               int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	EventID          int       `gorm:"column:event_id;type:int"`
	Name             string    `gorm:"column:name;size:100;not null"`
	ShortDescription string    `gorm:"column:short_description;size:255"`
	AvatarURL        string    `gorm:"column:avatar_url;size:255"`
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp;not null"`
}
