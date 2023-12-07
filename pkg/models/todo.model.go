package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ToDo struct {
	ID          string    `gorm:"type:char(36);primary_key;" json:"id"`
	Title       string    `gorm:"size:100;not null" json:"title"`
	Content     string    `gorm:"text;not null;" json:"content"`
	AuthorID    string    `gorm:"not null" json:"author_id"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null;ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	CompletedAt time.Time `gorm:"default:null" json:"completed_at"`
}

func (todo *ToDo) BeforeCreate(*gorm.DB) error {
	todo.ID = uuid.NewV4().String()
	return nil
}

// Function to return a simplified version of the todo
func (todo *ToDo) Simplified() TodoSimplified {
	return TodoSimplified{
		ID:          todo.ID,
		Title:       todo.Title,
		Content:     todo.Content,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	}
}

type TodoSimplified struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CompletedAt time.Time `json:"completed_at"`
}
