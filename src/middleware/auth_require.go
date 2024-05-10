package middleware

import (
	"errors"
	"net/http"
	"social-network/utils"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func AuthRequire() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := utils.VerifyTokenHeader(c)
		if err != nil {
			utils.WriteErrorResponse(c, http.StatusBadRequest, errors.New("token cannot verify"))
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.WriteErrorResponse(c, http.StatusBadRequest, errors.New("err"))
		}
		authData, ok := claims["auth"].(map[string]interface{})
		if !ok {
			utils.WriteErrorResponse(c, http.StatusBadRequest, errors.New("auth data not found"))
			return
		}
		c.Set("authData", authData)
		c.Next()
	}
}
