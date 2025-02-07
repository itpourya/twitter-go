package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"twitter-go-api/internal/pkg/jwt"
)

func Authorization(jwtService jwt.TokenService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authToken := context.GetHeader("Authorization")
		if authToken == "" {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Missing token", "status": false})
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing token", "status": false})
			return
		}
		user, err := jwtService.ValidateToken(authToken)
		if err != nil {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Missing token", "status": false})
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token", "status": false})
			return
		}

		context.Set("user_email", user.Email)
		context.Set("user_id", user.ID)
		context.Set("username", user.Username)

		return
	}
}
