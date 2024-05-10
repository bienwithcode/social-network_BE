package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin, Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token")

		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		c.Next()
	}
}
