package repository

import (
	"database/sql"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"my_fin/src/data_provider"
	"regexp"
	"time"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"email"`
	Password string `json:"password"`
	UserSign string `json:"user_sign"`
}

type RegisterUser struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	UserSign   string `json:"user_sign"`
}

type TokenData struct {
	AccessToken  string
	RefreshToken string
	ValidUntil   int64
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
		ur.db.InsertQuery("user_category", map[string]interface{}{
			"u_id":                u.ID,
			"categories":          "",
			"categories_incoming": "",
		})
		return
	}
	return u, true, nil
}

func (ur *UserRepository) ValidateUser(login string, password string) (u User, res bool) {
	//check login is valid mail
	if !reEmail.MatchString(login) {
		return
	}

	row := ur.db.SelectRow("SELECT user_id, login, password_hash FROM users WHERE login = ?", login)

	errU := row.Scan(&u.ID, &u.Username, &u.Password)
	if errU != nil && errU != sql.ErrNoRows {
		return
	}
	errC := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if errC == nil {
		return u, true
	}
	return u, false
}

/**
return access_token, refresh_token
*/
func (ur *UserRepository) CreateToken(userId uint64, userSign string) (*TokenData, error) {
	tData := &TokenData{
		ValidUntil: time.Now().Add(time.Minute * time.Duration(ur.jwtLiveTime)).Unix(),
	}
	atClaims := jwt.MapClaims{
		//"authorized": true,
		"user_id": userId,
		"exp":     tData.ValidUntil,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(ur.jwtKey))
	if err != nil {
		return tData, err
	}
	tData.AccessToken = token
	refresh, errR := ur.generateRefreshToken(userId, userSign)
	if errR != nil {
		return tData, errR
	}
	tData.RefreshToken = refresh
	return tData, nil
}

func (ur *UserRepository) generateRefreshToken(userId uint64, userSign string) (string, error) {
	ur.removeExpiredTokens(userId)
	nowTime := time.Now()
	uid := uuid.New()
	ur.db.InsertQuery("users_refresh_tokens", map[string]interface{}{
		"user_id":       userId,
		"refresh_token": uid,
		"fingerprint":   userSign,
		"created_at":    nowTime.Unix(),
		"expires_at":    nowTime.Unix() + 60*86400, //valid in 60 days
	})
	return uid.String(), nil
}

func (ur *UserRepository) ValidateToken(tokenString string) (userID interface{}, valid bool) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ur.jwtKey), nil
	})

	if err != nil {
		return
	}

	if token.Valid {
		for i, v := range token.Claims.(jwt.MapClaims) {
			if i == "user_id" {
				return v, true
			}
		}
	}
	return
}

/**
browser send fingerprint. validate pair token + fingerprint. If false - token is wrong or stolen
*/
func (ur *UserRepository) ValidateRefreshToken(user interface{}, refreshToken string, fingerPrint string) bool {
	sqlR := "SELECT refresh_token, fingerprint, expires_at FROM users_refresh_tokens WHERE user_id = ?"
	rows, errRf := ur.db.SelectQuery(sqlR, user)
	if errRf != nil {
		return false
	}
	if rows != nil {
		defer rows.Close()
	}

	curTime := time.Now()

	for rows.Next() {
		var rT string
		var fP string
		var expired int64
		errS := rows.Scan(&rT, &fP, &expired)
		if errS != nil {
			continue
		}

		if curTime.Unix() >= expired {
			ur.removeExpiredTokens(user)
			continue
		}
		if rT == refreshToken {
			return fP == fingerPrint
		}
	}
	return false
}

func (ur *UserRepository) removeExpiredTokens(userId interface{}) {
	nowTime := time.Now()
	_, _ = ur.db.Exec("DELETE FROM expires_at WHERE user_id = ? AND expires_at < ?", userId, nowTime.Unix())
}
