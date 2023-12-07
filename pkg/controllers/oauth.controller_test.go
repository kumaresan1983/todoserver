// controllers_test.go
package controllers

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/kumaresan1983/todoserver/pkg/initializers"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"golang.org/x/oauth2"
// )

// // GoogleOAuthConfig represents the Google OAuth configuration
// type GoogleOAuthConfig interface {
// 	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
// }

// // OAuthConfigWrapper is a wrapper type for *oauth2.Config that satisfies the GoogleOAuthConfig interface
// type OAuthConfigWrapper struct {
// 	config *oauth2.Config
// }

// func (o *OAuthConfigWrapper) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
// 	return o.config.Exchange(ctx, code)
// }

// // HTTPClient represents the HTTP client for fetching user data
// type HTTPClient interface {
// 	Get(url string) (*http.Response, error)
// }

// // MockGoogleOAuthConfig is a mock implementation of the GoogleOAuthConfig interface
// type MockGoogleOAuthConfig struct {
// 	mock.Mock
// }

// func (m *MockGoogleOAuthConfig) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
// 	args := m.Called(ctx, code)
// 	return args.Get(0).(*oauth2.Token), args.Error(1)
// }

// // MockHTTPClient is a mock implementation of the HTTPClient interface
// type MockHTTPClient struct {
// 	mock.Mock
// }

// func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
// 	args := m.Called(url)
// 	return args.Get(0).(*http.Response), args.Error(1)
// }

// var (
// 	googleOAuthConfig GoogleOAuthConfig
// 	httpClient        HTTPClient
// )

// func init() {
// 	// Set the initial implementations to the actual implementations
// 	googleOAuthConfig = &OAuthConfigWrapper{config: &oauth2.Config{}}
// 	httpClient = http.DefaultClient
// }

// // SetGoogleOAuthConfigProvider sets the Google OAuth configuration provider
// func SetGoogleOAuthConfigProvider(provider GoogleOAuthConfig) {
// 	googleOAuthConfig = provider
// }

// // SetHTTPClient sets the HTTP client
// func SetHTTPClient(client HTTPClient) {
// 	httpClient = client
// }

// func TestGoogleCallback(t *testing.T) {
// 	// Set up the Gin router
// 	router := gin.Default()
// 	router.GET("/auth/google/callback", GoogleCallback)

// 	// Set up the test database
// 	initializers.ConnectDB()
// 	// defer initializers.CloseDB()

// 	// Mock Google OAuth configuration
// 	mockOAuthConfig := new(MockGoogleOAuthConfig)
// 	mockOAuthConfig.On("Exchange", mock.Anything, "mock_code").Return(&oauth2.Token{AccessToken: "mock_access_token"}, nil)

// 	// Mock HTTP client for user data fetch
// 	mockHTTPClient := new(MockHTTPClient)
// 	mockResponse := &http.Response{
// 		StatusCode: http.StatusOK,
// 		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"id": "123", "email": "test@example.com", "name": "Test User"}`))),
// 	}
// 	mockHTTPClient.On("Get", "https://www.googleapis.com/oauth2/v2/userinfo?access_token=mock_access_token").Return(mockResponse, nil)

// 	// Replace the actual implementations with mocks
// 	SetGoogleOAuthConfigProvider(mockOAuthConfig)
// 	SetHTTPClient(mockHTTPClient)

// 	// Perform a GET request to /auth/google/callback with mock parameters
// 	state := "randomstate"
// 	code := "mock_code"
// 	url := "/auth/google/callback?state=" + state + "&code=" + code
// 	req, err := http.NewRequest("GET", url, nil)
// 	assert.NoError(t, err)

// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assert the HTTP status code
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Check if the JWT token was generated
// 	var responseMap map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &responseMap)
// 	assert.NoError(t, err)

// 	assert.Contains(t, responseMap, "token")
// 	tokenValue, ok := responseMap["token"].(string)
// 	assert.True(t, ok)
// 	assert.NotEmpty(t, tokenValue)

// 	// // Check if log messages were generated
// 	// logEntries := logrus.StandardLogger().Entries
// 	// assert.NotEmpty(t, logEntries)
// 	// assert.Contains(t, logEntries[0].Message, "Google user data")
// 	// assert.Contains(t, logEntries[1].Message, "Google login successful")

// 	// Assert that the expectations were met
// 	mockOAuthConfig.AssertExpectations(t)
// 	mockHTTPClient.AssertExpectations(t)
// }
