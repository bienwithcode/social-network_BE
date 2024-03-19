package utils

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

const JWT_ENV = "JWT_SECRET"

type MetaToken struct {
	Email     string
	ExpiredAt time.Time
}

type AccessToken struct {
	Claims MetaToken
}

func IssueToken(Data map[string]interface{}, ExpiredAt time.Duration) (string, int64, error) {

	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()

	jwtSecretKey := GodotEnv(JWT_ENV)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return accessToken, expiredAt, err
	}

	return accessToken, expiredAt, nil
}

func VerifyTokenHeader(ctx *gin.Context) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	jwtSecretKey := GodotEnv(JWT_ENV)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func VerifyToken(accessToken string) (*jwt.Token, error) {
	jwtSecretKey := GodotEnv(JWT_ENV)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}
