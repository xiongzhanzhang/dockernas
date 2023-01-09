package middleware

import (
	"strings"
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		token := c.Request.Header.Get("token")
		// log.Println(token)
		// log.Println(path)

		if path == "/api/network/ipv4" || path == "/api/network/ipv6" {
			if strings.Index(c.Request.Host, "localhost") != 0 || strings.Index(c.Request.RemoteAddr, "127.0.0.1") != 0 {
				c.String(404, "")
				c.Abort()
			}
			c.Next()
		} else if path == "/api/login" {
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
