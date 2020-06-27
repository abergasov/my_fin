package repository

import (
	"database/sql"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"my_fin/src/data_provider"
	"regexp"
	"time"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"email"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

var reEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type UserRepository struct {
	jwtKey      string
	jwtLiveTime int64
	db          *data_provider.DBAdapter
}

func InitUserRepository(db *data_provider.DBAdapter, jwtKey string, jwtLive int64) *UserRepository {
	return &UserRepository{
		jwtKey:      jwtKey,
		jwtLiveTime: jwtLive,
		db:          db,
	}
}

func (ur *UserRepository) RegisterUser(rU *RegisterUser) (u User, exist bool, err error) {
	if rU.Password != rU.RePassword {
		return
	}
	//check login is valid mail
	if !reEmail.MatchString(rU.Email) {
		return u, false, errors.New("42")
	}

	//check mail already exist
	row := ur.db.SelectRow("SELECT user_id FROM users WHERE login = ?", rU.Email)
	errU := row.Scan(&u.ID)
	if errU != nil && errU != sql.ErrNoRows {
		return u, false, errors.New("42")
	}
	if errU != nil && errU == sql.ErrNoRows {
		passwordHash, errP := bcrypt.GenerateFromPassword([]byte(rU.Password), 8)
		if errP != nil {
			return u, false, errors.New("42")
		}
		u.ID = uint64(ur.db.InsertQuery("users", map[string]interface{}{"login": rU.Email, "password_hash": passwordHash}))
		return
	}
	return u, true, nil
}

func (ur *UserRepository) ValidateUser(login string, password string) (u User, res bool) {
	//check login is valid mail
	if !reEmail.MatchString(login) {
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return
	}
	sqlU := "SELECT user_id, login, password_hash FROM users WHERE login = ? AND password_hash = ?"
	row := ur.db.SelectRow(sqlU, login, passwordHash)

	errU := row.Scan(&u.ID, &u.Username, &u.Password)
	if errU != nil && err != sql.ErrNoRows {
		return
	}
	return u, true
}

func (ur *UserRepository) CreateToken(userId uint64) (string, error) {
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        time.Now().Add(time.Minute * time.Duration(ur.jwtLiveTime)).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(ur.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
