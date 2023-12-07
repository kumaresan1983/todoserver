package initializers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// Set up a temporary environment variable for testing
	os.Setenv("DatabaseUserName", "your_username")
	os.Setenv("DatabasePassword", "your_password")
	os.Setenv("DatabaseHost", "localhost")
	os.Setenv("DatabasePort", "3306")
	os.Setenv("DatabaseName", "your_database_name")

	// Call the ConnectDB function
	err := ConnectDB()

	// Assertions
	assert.NoError(t, err, "ConnectDB should not return an error")
	assert.NotNil(t, DB, "DB should be initialized")

	// Clean up the temporary environment variables
	os.Unsetenv("DatabaseUserName")
	os.Unsetenv("DatabasePassword")
	os.Unsetenv("DatabaseHost")
	os.Unsetenv("DatabasePort")
	os.Unsetenv("DatabaseName")
}
