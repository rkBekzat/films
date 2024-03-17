package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rkBekzat/films/internal/model"
	"github.com/rkBekzat/films/internal/repository"
)

const (
	salt       = "desgerfr3241rfgfqwer"
	signingKey = "dqwfdqwdewvafsdvfcs12erf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type auth struct {
	repo repository.Account
}

func NewAuthorization(repo repository.Account) Account {
	return &auth{repo: repo}
}

func (a *auth) CreateUser(user *model.User) error {
	if user.Gender != "male" && user.Gender != "female" {
		return fmt.Errorf("the gender doesn't exist. male or female")
	}
	exist, err := a.repo.EmailExist(user.Email)
	if err != nil {
		return fmt.Errorf("failed to check email: %s", err.Error())
	}
	if exist {
		return fmt.Errorf("the email %s exist", user.Email)
	}
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *auth) GenerateToken(username, password string) (string, error) {

	user, err := a.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (a *auth) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
