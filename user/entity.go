package user

import (
	"time"
)

type User struct {
	ID              int     `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name            string    `gorm:"column:name;size:100;not null"`
	Email           string    `gorm:"column:email;size:255;unique;not null"`
	Phone           string    `gorm:"column:phone;size:25"`
	AvatarFile      string    `gorm:"column:avatar_file;size:255"`
	PasswordHash    string    `gorm:"column:password_hash;size:255;not null"`
	Role            string    `gorm:"column:role;size:10"`
	Token           string    `gorm:"column:token;size:255"`
	IsEmailVerified int32     `gorm:"column:is_email_verified;type:int"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;not null"`
}
