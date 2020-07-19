package routes

import (
	"my_fin/backend/pkg/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	RefreshCookie = "rc"
	TokenCookie   = "tc"
)

func (ar *AppRouter) Login(c *gin.Context) {
	var u repository.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}

	// compare the user from the request, with the one we defined:
	uR, valid := ar.userRepo.ValidateUser(u.Username, u.Password)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}

	tData, err := ar.userRepo.CreateToken(uR.ID, u.UserSign)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}

	uR.Password = ""
	uR.UserSign = ""
	ar.setSecretCookie(c, TokenCookie, tData.AccessToken)
	ar.setSecretCookie(c, RefreshCookie, tData.RefreshToken)
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": tData.AccessToken, "user": uR})
}

func (ar *AppRouter) Register(c *gin.Context) {
	var u repository.RegisterUser
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}
	user, userExist, errR := ar.userRepo.RegisterUser(&u)
	if errR != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": errR.Error()})
		return
	}
	if userExist {
		c.JSON(http.StatusConflict, gin.H{"ok": false, "error": "User exist"})
		return
	}

	tData, err := ar.userRepo.CreateToken(user.ID, u.UserSign)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Error while register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "token": tData.AccessToken, "user": user})
}

func (ar *AppRouter) Refresh(c *gin.Context) {

}

func (ar *AppRouter) Logout(c *gin.Context) {
	ar.setSecretCookie(c, TokenCookie, "")
	ar.setSecretCookie(c, RefreshCookie, "")
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (ar *AppRouter) setSecretCookie(c *gin.Context, keyName, keyValue string) {
	liveTime := int(ar.config.JWTLive) * 60
	path := "/api/data"
	if keyName == RefreshCookie {
		liveTime = 60 * 86400
		path = "/api/auth/refresh"
	}
	c.SetCookie(keyName, keyValue, liveTime, path, ar.config.AppDomain, ar.config.SSLEnable, true)
}
