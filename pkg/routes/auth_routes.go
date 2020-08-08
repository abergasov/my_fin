package routes

import (
	"my_fin/backend/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	RefreshCookie     = "rc"
	TokenCookie       = "tc"
	UserIDCookie      = "u"
	FingerPrintHeader = "m"
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
	ar.setSecretCookie(c, UserIDCookie, strconv.FormatInt(uR.ID, 10))
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

	ar.setSecretCookie(c, TokenCookie, tData.AccessToken)
	ar.setSecretCookie(c, RefreshCookie, tData.RefreshToken)
	ar.setSecretCookie(c, UserIDCookie, strconv.FormatInt(user.ID, 10))
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": tData.AccessToken, "user": user})
}

func (ar *AppRouter) Refresh(c *gin.Context) {
	usrFgp := c.GetHeader(FingerPrintHeader)
	rTokenC, errR := c.Cookie(RefreshCookie)
	if errR != nil {
		return
	}
	uID, err := c.Cookie(UserIDCookie)
	if err != nil {
		return
	}
	res := ar.userRepo.ValidateRefreshToken(uID, rTokenC, usrFgp)
	if res {
		uIDInt, errCnv := strconv.ParseInt(uID, 10, 64)
		if errCnv != nil {
			return
		}
		tData, errCr := ar.userRepo.CreateToken(uIDInt, rTokenC)
		if errCr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
			return
		}

		ar.setSecretCookie(c, TokenCookie, tData.AccessToken)
	}
	return
}

func (ar *AppRouter) Logout(c *gin.Context) {
	ar.setSecretCookie(c, TokenCookie, "")
	ar.setSecretCookie(c, RefreshCookie, "")
	ar.setSecretCookie(c, UserIDCookie, "")
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (ar *AppRouter) setSecretCookie(c *gin.Context, keyName, keyValue string) {
	liveTime := ar.userRepo.GetTokenValidUntil()
	path := "/api/data"
	if keyName == RefreshCookie || keyName == UserIDCookie {
		liveTime = 60 * 86400
		path = "/api/auth/refresh"
	}
	c.SetCookie(keyName, keyValue, int(liveTime), path, ar.config.AppDomain, ar.config.SSLEnable, true)
}
