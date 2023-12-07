package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_SetID(t *testing.T) {
	user := Users{}
	user.SetID()

	assert.NotEmpty(t, user.ID, "ID should be set")
}

func TestUsers_Validate_ValidData(t *testing.T) {
	user := Users{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	err := user.Validate()

	assert.Nil(t, err, "Validation should pass for valid data")
}

func TestUsers_Validate_InvalidData(t *testing.T) {
	user := Users{
		Email: "john@example.com",
	}

	err := user.Validate()

	assert.NotNil(t, err, "Validation should fail for invalid data")
}

func TestUsers_BeforeCreate(t *testing.T) {
	user := Users{}
	err := user.BeforeCreate(nil)

	assert.Nil(t, err, "BeforeCreate should not return an error")
	assert.NotEmpty(t, user.ID, "BeforeCreate should set ID")
}

// Add more test cases as needed to cover other functions and branches in your code.
