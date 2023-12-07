package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kumaresan1983/todoserver/pkg/initializers"
	"github.com/kumaresan1983/todoserver/pkg/models"
	"github.com/kumaresan1983/todoserver/pkg/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const projectDirName = "todoserver" // change to relevant project name

type GoogleUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

func GoogleLogin(c *gin.Context) {

	// Assuming you have a function to create the Google OAuth2 configuration
	googleConfig := createGoogleOAuthConfig()

	// Construct the URL for Google login
	url := googleConfig.AuthCodeURL("randomstate")

	// Redirect to the Google login URL
	c.Redirect(http.StatusSeeOther, url)
}

func GoogleCallback(c *gin.Context) {

	var state string = "/"

	if c.Query("state") != "" {
		state = c.Query("state")
	}

	// state := c.Query("state")
	if state != "randomstate" {
		c.String(http.StatusBadRequest, "States don't Match!!")
		return
	}

	code := c.Query("code")

	// Assuming you have a function to create the Google OAuth2 configuration
	googleConfig := createGoogleOAuthConfig()

	// Exchange the code for a token
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Code-Token Exchange Failed")
		return
	}

	// Fetch user data using the access token
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.String(http.StatusInternalServerError, "User Data Fetch Failed")
		return
	}
	defer resp.Body.Close()

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, resp.Body)
	if err != nil {
		return
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return
	}

	google_user := &GoogleUserResult{
		Id:             GoogleUserRes["id"].(string),
		Email:          GoogleUserRes["email"].(string),
		Verified_email: GoogleUserRes["verified_email"].(bool),
		Name:           GoogleUserRes["name"].(string),
		Given_name:     GoogleUserRes["given_name"].(string),
		Picture:        GoogleUserRes["picture"].(string),
		Locale:         GoogleUserRes["locale"].(string),
	}

	now := time.Now()
	email := strings.ToLower(google_user.Email)

	user_data := models.Users{
		Name:      google_user.Name,
		Email:     email,
		Provider:  "Google",
		UpdatedAt: now,
	}

	fmt.Printf("%+v", user_data)

	if initializers.DB.Model(&user_data).Where("email = ?", email).Updates(&user_data).RowsAffected == 0 {
		initializers.DB.Create(&user_data)
	}

	tokens, err := utils.GenerateJWT(user_data.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokens})

}

// Create Google OAuth2 configuration function
func createGoogleOAuthConfig() *oauth2.Config {

	// projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	// currentWorkDirectory, _ := os.Getwd()
	// rootPath := projectName.Find([]byte(currentWorkDirectory))

	config, _ := initializers.LoadConfig(".")

	return &oauth2.Config{
		ClientID:     config.GoogleClientID,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  config.GoogleOAuthRedirectUrl,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
}
