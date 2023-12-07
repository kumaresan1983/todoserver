package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string    `gorm:"type:char(36);primary_key;" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null;" json:"email"`
	Provider  string    `gorm:"type:varchar(100);default:'local';"`
	Todos     []ToDo    `gorm:"foreignKey:AuthorID" json:"todos,omitempty"`
	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;ON UPDATE CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (user *Users) BeforeCreate(*gorm.DB) error {
	user.SetID()
	return nil
}

// SetID generates a new UUID and sets it as the user's ID
func (user *Users) SetID() {
	user.ID = uuid.NewV4().String()
}

// Validate checks if the user data is valid
func (user *Users) Validate() error {
	if user.Name == "" {
		return gorm.ErrRecordNotFound
	}
	return nil
}
