package auth_helper

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"regexp"
	"strings"
	"time"
)

const (
	Salt       = "hjqrhjqw124617ajfhajs"
	SigningKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	TokenTTL   = 7 * 24 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

var basicRegex = regexp.MustCompile(`^Basic (.+)$`)

var bearerRegex = regexp.MustCompile(`^Bearer (.+)$`)

func ParseHttpHeaderValue(value string) (Authorization, error) {
	matchesBasic := basicRegex.FindStringSubmatch(value)
	if len(matchesBasic) > 0 {
		return ParseBase(matchesBasic[1])
	}

	matchesBearer := bearerRegex.FindStringSubmatch(value)
	if len(matchesBearer) > 0 {
		return TokenAuth{Token: matchesBearer[1]}, nil
	}

	return nil, errors.New("Unsupported authorization type: " + value)
}

func ParseBase(encoded string) (Authorization, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	decodedStr := string(decoded)
	parts := strings.Split(decodedStr, ":")
	if len(parts) != 2 {
		return nil, errors.New("<4f95c0c3> Invalid basic auth string")
	}

	return &CredentialsAuth{
		Login:    parts[0],
		Password: parts[1],
	}, nil
}

func (ca CredentialsAuth) ToHttpHeaderValue() string {

	credentials := fmt.Sprintf("%s:%s", ca.Login, ca.Password)
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	return fmt.Sprintf("Basic %s", encoded)

}

func (ta TokenAuth) ToHttpHeaderValue() string {
	return fmt.Sprintf("Bearer %s", ta.Token)
}

// Функция String для структуры CredentialsAuth
func (ca CredentialsAuth) String() string {
	return fmt.Sprintf("%s:%s", ca.Login, ca.Password)
}

// Функция String для структуры TokenAuth
func (ta TokenAuth) String() string {
	return ta.Token
}

func NewTokenAuth(token string) Authorization {
	return TokenAuth{token}
}

func NewCredentialsAuth(login, password string) Authorization {
	return &CredentialsAuth{login, password}
}

func (ta TokenAuth) GetUserIdFromToken() (string, error) {
	token, err := jwt.ParseWithClaims(ta.Token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
