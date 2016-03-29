package middlewares

import "github.com/gin-gonic/gin"

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") == "test" {
			c.Next()
		} else {
			c.Status(403)
			c.Abort()
		}
	}
}
