package routes

import (
	"github.com/gin-gonic/gin"
	"my_fin/src/repository"
	"net/http"
)

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

	token, err := ar.userRepository.CreateToken(uR.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid login/password"})
		return
	}
	c.SetCookie("token", token, int(ar.config.JWTLive), "/", ar.config.AppDomain, ar.config.SSLEnable, true)
	c.JSON(http.StatusOK, gin.H{"ok": true})
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

	token, err := ar.userRepository.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Error while register"})
		return
	}
	c.SetCookie("token", token, int(ar.config.JWTLive), "/", ar.config.AppDomain, ar.config.SSLEnable, true)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
