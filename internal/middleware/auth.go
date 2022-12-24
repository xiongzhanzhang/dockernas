package middleware

import (
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		token := c.Request.Header.Get("token")
		// log.Println(token)
		// log.Println(path)

		if path == "/api/login" {
			c.Next()
		} else {
			if service.IsTokenValid(token) == false {
				c.JSON(555, gin.H{"msg": "Authentication error"})
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
