package middleware

import (
	"github.com/gin-gonic/gin"
	"my_fin/backend/pkg/repository"
	"net/http"
)

const tokenCookie = "a"
const refreshCookie = "z"
const fingerPrintHeader = "m"

func AuthMiddleware(userRepo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(tokenCookie)
		if err != nil || len(token) < 10 {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
			c.Abort()
			return
		}

		uId, valid := userRepo.ValidateToken(token)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
			c.Abort()
			return
		}

		c.Set("user_id", uId)
		c.Next()
	}
}
