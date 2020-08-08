package repository

import (
	"database/sql"
	"errors"
	"my_fin/backend/pkg/database"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"email"`
	Password string `json:"password,omitempty"`
	UserSign string `json:"user_sign,omitempty"`

	MandatoryPercent int64 `json:"mandatory_percent"`
	LivePercent      int64 `json:"live_percent"`
	BlackDayPercent  int64 `json:"black_day_percent"`
	InvestPercent    int64 `json:"invest_percent"`
	SpendingPercent  int64 `json:"spending_percent"`
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
	db          *database.DBAdapter
}

func InitUserRepository(db *database.DBAdapter, jwtKey string, jwtLive int64) *UserRepository {
	return &UserRepository{
		jwtKey:      jwtKey,
		jwtLiveTime: jwtLive,
		db:          db,
	}
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

func (ur *UserRepository) ValidateUser(login, password string) (u User, res bool) {
	// check login is valid mail
	if !reEmail.MatchString(login) {
		return
	}

	row := ur.db.SelectRow("SELECT user_id, login, password_hash, mandatory_percent, live_percent, black_day_percent, invest_percent, spending_percent FROM users WHERE login = ?", login)

	errU := row.Scan(&u.ID, &u.Username, &u.Password, &u.MandatoryPercent, &u.LivePercent, &u.BlackDayPercent, &u.InvestPercent, &u.SpendingPercent)
	if errU != nil && errU != sql.ErrNoRows {
		return
	}
	errC := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return u, errC == nil
}

/**
return access_token, refresh_token
*/
func (ur *UserRepository) CreateToken(userID int64, userSign string) (*TokenData, error) {
	tData := &TokenData{
		ValidUntil: ur.GetTokenValidUntil(),
	}
	atClaims := jwt.MapClaims{
		// "authorized": true,
		"user_id": userID,
		"exp":     tData.ValidUntil,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(ur.jwtKey))
	if err != nil {
		return tData, err
	}
	tData.AccessToken = token
	refresh, errR := ur.generateRefreshToken(userID, userSign)
	if errR != nil {
		return tData, errR
	}
	tData.RefreshToken = refresh
	return tData, nil
}

func (ur *UserRepository) generateRefreshToken(userID int64, userSign string) (string, error) {
	ur.removeExpiredTokens(userID)
	nowTime := time.Now()
	uid := uuid.New()
	iserted := ur.db.InsertQuery("users_refresh_tokens", map[string]interface{}{
		"user_id":       userID,
		"refresh_token": uid,
		"fingerprint":   userSign,
		"created_at":    nowTime.Unix(),
		"expires_at":    nowTime.Unix() + 60*86400, // valid in 60 days
	})
	if iserted > 0 {
		return uid.String(), nil
	}
	return "", errors.New("error set refresh token")
}

func (ur *UserRepository) removeExpiredTokens(userID interface{}) {
	nowTime := time.Now()
	_, _ = ur.db.Exec("DELETE FROM expires_at WHERE user_id = ? AND expires_at < ?", userID, nowTime.Unix())
}

func (ur *UserRepository) RegisterUser(rU *RegisterUser) (u User, exist bool, err error) {
	if rU.Password != rU.RePassword {
		return
	}
	// check login is valid mail
	if !reEmail.MatchString(rU.Email) {
		return u, false, errors.New("42")
	}

	// check mail already exist
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
		u.ID = ur.db.InsertQuery("users", map[string]interface{}{"login": rU.Email, "password_hash": passwordHash})

		u.MandatoryPercent = 30
		u.LivePercent = 20
		u.BlackDayPercent = 15
		u.InvestPercent = 15
		u.SpendingPercent = 20

		ur.db.InsertQuery("user_category", map[string]interface{}{
			"u_id":                u.ID,
			"categories":          "",
			"categories_incoming": "",
		})
		return
	}
	return u, true, nil
}

/**
browser send fingerprint. validate pair token + fingerprint. If false - token is wrong or stolen
*/
func (ur *UserRepository) ValidateRefreshToken(user interface{}, refreshToken, fingerPrint string) bool {
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

func (ur *UserRepository) GetTokenValidUntil() int64 {
	return time.Now().Add(time.Minute * time.Duration(ur.jwtLiveTime)).Unix()
}
