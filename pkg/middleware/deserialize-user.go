package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumaresan1983/todoserver/pkg/initializers"
	"github.com/kumaresan1983/todoserver/pkg/models"
	"github.com/kumaresan1983/todoserver/pkg/utils"
	"github.com/sirupsen/logrus"
)

// Auth is a middleware for handling authentication using JWT.
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			logrus.Warn("Request does not contain an access token")
			context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		email, err := utils.ValidateJWTToken(tokenString)
		if err != nil {
			logrus.WithError(err).Warn("Failed to validate JWT token")
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		logrus.Infof("Token validation successful for email: %s", email)

		var user models.Users
		result := initializers.DB.First(&user, "email = ?", email)
		if result.Error != nil {
			logrus.WithError(result.Error).Warn("User not found for the provided token")
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no longer exists"})
			return
		}

		logrus.Infof("User found: ID=%s, Name=%s", user.ID, user.Name)

		context.Set("currentUser", user)
		context.Next()
	}
}
