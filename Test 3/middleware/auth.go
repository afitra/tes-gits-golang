package Middleware

import (
	"gits/auth"
	"gits/helper"
	"gits/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func IsLogin(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("token")
		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, errorValidationToken := token.Claims.(jwt.MapClaims)
		if !errorValidationToken || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {

			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
