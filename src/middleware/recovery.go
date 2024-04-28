package middleware

import (
	"net/http"
	"social-network/utils"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.WriteErrorResponse(c, http.StatusBadRequest, r)
				c.Abort()
			}
		}()
		c.Next()
	}
}
