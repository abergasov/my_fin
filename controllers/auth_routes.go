package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

func (ar *AppRouter) Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := ar.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
}
func (ar *AppRouter) CreateToken(userId uint64) (string, error) {
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        time.Now().Add(time.Minute * time.Duration(ar.config.JWTLive)).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(ar.config.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
