package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kumaresan1983/todoserver/pkg/initializers"
	"github.com/kumaresan1983/todoserver/pkg/models"
	"github.com/kumaresan1983/todoserver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, func()) {
	// Set up an SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	// Run migrations
	err = db.AutoMigrate(&models.Users{}, &models.ToDo{})
	if err != nil {
		panic("Failed to run migrations on test database")
	}

	return db, func() {
		// Clean up by closing the database connection
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}

func TestAuthMiddleware(t *testing.T) {
	// Set up the test database
	testDB, cleanup := setupTestDB()
	defer cleanup()

	// Use the test database for the application
	initializers.DB = testDB

	// Create a new Gin router
	router := gin.New()

	// Set up the Auth middleware
	router.Use(Auth())

	// Define a test route that uses the middleware
	router.GET("/test", func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		assert.True(t, exists, "currentUser should exist in the context")
		assert.IsType(t, &models.Users{}, user, "currentUser should be of type Users")
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Test case 1: Missing Authorization header
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "request does not contain an access token")

	// Test case 2: Invalid JWT token
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "invalid_token")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid number of segments")

	// Test case 3: Valid JWT token and user found
	validToken, err := utils.GenerateJWT("test@example.com")
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "the user belonging to this token no longer exists")
}
