package routes

import (
	"github.com/gin-gonic/gin"
	"my_fin/src/repository"
	"net/http"
)

const tokenCookie = "a"
const refreshCookie = "z"
const fingerPrintHeader = "m"

func (ar *AppRouter) Login(c *gin.Context) {
	var u repository.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}

	//compare the user from the request, with the one we defined:
	uR, valid := ar.userRepository.ValidateUser(u.Username, u.Password)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}

	tData, err := ar.userRepository.CreateToken(uR.ID, u.UserSign)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}
	ar.setSecretCookie(c, tokenCookie, tData.AccessToken)
	ar.setSecretCookie(c, refreshCookie, tData.RefreshToken)
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": tData.AccessToken})
}

func (ar *AppRouter) Register(c *gin.Context) {
	var u repository.RegisterUser
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}
	user, userExist, errR := ar.userRepository.RegisterUser(&u)
	if errR != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": errR.Error()})
		return
	}
	if userExist {
		c.JSON(http.StatusConflict, gin.H{"ok": false, "error": "User exist"})
		return
	}

	tData, err := ar.userRepository.CreateToken(user.ID, u.UserSign)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Error while register"})
		return
	}
	ar.setSecretCookie(c, tokenCookie, tData.AccessToken)
	ar.setSecretCookie(c, refreshCookie, tData.RefreshToken)
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": tData.AccessToken})
}

func (ar *AppRouter) Logout(c *gin.Context) {
	ar.setSecretCookie(c, tokenCookie, "")
	ar.setSecretCookie(c, refreshCookie, "")
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (ar *AppRouter) Refresh(c *gin.Context) {
	tokenRefresh, errRf := c.Cookie(refreshCookie)
	if errRf != nil || len(tokenRefresh) < 10 {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false})
		return
	}

	var t struct {
		UserId int64 `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false})
		return
	}

	if ar.userRepository.ValidateRefreshToken(t.UserId, tokenRefresh, c.GetHeader(fingerPrintHeader)) {
		tData, err := ar.userRepository.CreateToken(uint64(t.UserId), c.GetHeader(fingerPrintHeader))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false})
			return
		}
		ar.setSecretCookie(c, tokenCookie, tData.AccessToken)
		c.JSON(http.StatusOK, gin.H{"ok": true, "token": tData.AccessToken})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"ok": false})
	return
}

func (ar *AppRouter) setSecretCookie(c *gin.Context, keyName string, keyValue string) {
	liveTime := int(ar.config.JWTLive) * 60
	path := "/api/data"
	if keyName == refreshCookie {
		liveTime = 60 * 86400
		path = "/api/auth/refresh"
	}
	c.SetCookie(keyName, keyValue, liveTime, path, ar.config.AppDomain, ar.config.SSLEnable, true)
}

func (ar *AppRouter) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(tokenCookie)
		if err != nil || len(token) < 10 {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
			c.Abort()
			return
		}

		uId, valid := ar.userRepository.ValidateToken(token)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
			c.Abort()
			return
		}

		c.Set("user_id", uId)
		c.Next()
	}
}
