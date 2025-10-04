package middleware

import (
	"net/http"
	"strings"

	"github.com/Proudprogamer/goAuth/http/utils"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization :=c.GetHeader("Authorization")

		if authorization == ""{
			c.JSON(500, gin.H{
				"message" : "Authorization header not set",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Invalid auth header format",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")

		claims, err := utils.ValidateToken(token)

		if err!=nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Invalid or expired token",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("username", claims.Name)
		c.Set("email", claims.Email)

		c.Next()
	}

}