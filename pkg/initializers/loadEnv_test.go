// loadEnv_test.go
package initializers

import (
	"io/ioutil"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set up test environment (e.g., load configuration)
	loadConfigForTest(t)
	defer viper.Reset()

	// Perform assertions or additional testing logic
	config, err := LoadConfig(".")
	assert.NoError(t, err, "Error loading config")

	// Validate the loaded configuration values
	assert.Equal(t, "532663944661-u91cki5nuo08dkbbvu9os5bb8t2ae9es.apps.googleusercontent.com", config.GoogleClientID)
	assert.Equal(t, "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", config.GoogleClientSecret)
	assert.Equal(t, "http://localhost:8080/v1/api/auth/google/callback", config.GoogleOAuthRedirectUrl)
	assert.Equal(t, 60, config.TokenMaxAge)
	assert.Equal(t, "my_ultra_secure_secret", config.JWTTokenSecret)
	// ... (add other assertions for configuration values)
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	// Set up test environment with a non-existent directory
	nonExistentDir := "/non-existent-directory"
	viper.AddConfigPath(nonExistentDir)

	// Perform assertions for the expected behavior
	_, err := LoadConfig(".")
	assert.Error(t, err, "Not Found")
	assert.Contains(t, err.Error(), "Config File ")
	// Add more assertions as needed
}

func TestLoadConfig_UnmarshalError(t *testing.T) {
	// Set up test environment with a directory containing an invalid config file
	invalidConfigDir := "/invalid-config-directory"
	viper.AddConfigPath(invalidConfigDir)

	// Create an invalid configuration file
	invalidConfigFilePath := invalidConfigDir + "/local.env"
	invalidConfigFileContent := `invalid_yaml_content`
	err := ioutil.WriteFile(invalidConfigFilePath, []byte(invalidConfigFileContent), 0644)
	assert.Error(t, err, "system cannot find the path specified")

	// Perform assertions for the expected behavior
	_, err = LoadConfig(".")
	assert.Error(t, err, "Expected error for unmarshal failure")
	assert.Contains(t, err.Error(), "Not Found")
	// Add more assertions as needed
}

func TestLoadConfig_ReadConfigError(t *testing.T) {
	// Mock an error reading the configuration
	viperReadInConfigOrig := viperReadInConfig
	defer func() { viperReadInConfig = viperReadInConfigOrig }()
	viperReadInConfig = func() error {
		return assert.AnError
	}

	// Perform assertions for the expected behavior
	_, err := LoadConfig(".")
	assert.Error(t, err, "Expected error for read config failure")
	assert.Contains(t, err.Error(), "Not Found")
	// Add more assertions as needed
}

var viperReadInConfig = viper.ReadInConfig

func loadConfigForTest(t *testing.T) {
	// Set up a temporary directory for testing
	tmpDir := t.TempDir()

	// Create a sample configuration file in the temporary directory
	configFileContent := `
	JWT_SECRET=your_jwt_secret
	TOKEN_EXPIRED_IN=1h
	TOKEN_MAXAGE=3600
	GOOGLE_OAUTH_CLIENT_ID=your_google_client_id
	GOOGLE_OAUTH_CLIENT_SECRET=your_google_client_secret
	GOOGLE_OAUTH_REDIRECT_URL=http://example.com/auth/google/callback
	`
	configFilePath := tmpDir + "/local.env"
	err := ioutil.WriteFile(configFilePath, []byte(configFileContent), 0644)
	assert.NoError(t, err, "Error creating sample config file")

	// Load the configuration from the temporary directory
	viper.AddConfigPath(tmpDir)
	viper.SetConfigType("env")
	viper.SetConfigName("local")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	assert.NoError(t, err, "Error reading sample config file")
}

// // loadEnv_test.go
// package initializers

// import (
// 	"io/ioutil"
// 	"log"
// 	"testing"

// 	"github.com/spf13/viper"
// 	"github.com/stretchr/testify/assert"
// )

// func TestLoadConfig(t *testing.T) {
// 	// Set up test environment (e.g., load configuration)
// 	loadConfigForTest(t)
// 	defer viper.Reset()

// 	// Perform assertions or additional testing logic
// 	config, err := LoadConfig(".")
// 	assert.NoError(t, err, "Error loading config")

// 	// Validate the loaded configuration values
// 	assert.Equal(t, "http://example.com", config.FrontEndOrigin)
// 	// ... (add other assertions for configuration values)
// }

// func loadConfigForTest(t *testing.T) {
// 	// Set up a temporary directory for testing
// 	tmpDir := t.TempDir()

// 	// Create a sample configuration file in the temporary directory
// 	configFileContent := `
// 	FRONTEND_ORIGIN=http://example.com
// 	JWT_SECRET=your_jwt_secret
// 	TOKEN_EXPIRED_IN=1h
// 	TOKEN_MAXAGE=3600
// 	GOOGLE_OAUTH_CLIENT_ID=your_google_client_id
// 	GOOGLE_OAUTH_CLIENT_SECRET=your_google_client_secret
// 	GOOGLE_OAUTH_REDIRECT_URL=http://example.com/auth/google/callback
// 	GITHUB_OAUTH_CLIENT_ID=your_github_client_id
// 	GITHUB_OAUTH_CLIENT_SECRET=your_github_client_secret
// 	GITHUB_OAUTH_REDIRECT_URL=http://example.com/auth/github/callback
// 	`
// 	configFilePath := tmpDir + "/local.env"
// 	err := ioutil.WriteFile(configFilePath, []byte(configFileContent), 0644)
// 	if err != nil {
// 		log.Fatal("Error creating sample config file: ", err)
// 	}

// 	// Load the configuration from the temporary directory
// 	viper.AddConfigPath(tmpDir)
// 	viper.SetConfigType("env")
// 	viper.SetConfigName("local")
// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatal("Error reading sample config file: ", err)
// 	}
// }
