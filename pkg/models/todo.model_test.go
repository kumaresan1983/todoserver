package models

import (
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

func TestToDoModel(t *testing.T) {
	t.Run("Test BeforeCreate", func(t *testing.T) {
		todo := &ToDo{}
		db := &gorm.DB{} // Mock gorm.DB

		err := todo.BeforeCreate(db)

		assert.NoError(t, err, "BeforeCreate should not return an error")
		assert.NotEmpty(t, todo.ID, "BeforeCreate should set a non-empty ID")
	})

	t.Run("Test TodoSimplified", func(t *testing.T) {
		now := time.Now()
		todo := &ToDo{
			ID:          "123",
			Title:       "Test Title",
			Content:     "Test Content",
			Completed:   true,
			CreatedAt:   now,
			UpdatedAt:   now,
			CompletedAt: now,
		}

		simplified := todo.Simplified()

		assert.Equal(t, "123", simplified.ID)
		assert.Equal(t, "Test Title", simplified.Title)
		assert.Equal(t, "Test Content", simplified.Content)
		assert.True(t, simplified.Completed)
		assert.Equal(t, now, simplified.CreatedAt)
		assert.Equal(t, now, simplified.UpdatedAt)
		assert.Equal(t, now, simplified.CompletedAt)
	})
}
