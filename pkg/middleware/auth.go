package middleware

import (
	"my_fin/backend/pkg/logger"
	"my_fin/backend/pkg/repository"
	"my_fin/backend/pkg/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(routes.TokenCookie)
		if err != nil || len(token) < 10 {
			logger.Warning("auth error", "missed cookie")
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
			c.Abort()
			return
		}

		uID, valid := userRepo.ValidateToken(token)
		if !valid {
			logger.Warning("auth error", "invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
			c.Abort()
			return
		}

		c.Set("user_id", uID)
		c.Next()
	}
}
