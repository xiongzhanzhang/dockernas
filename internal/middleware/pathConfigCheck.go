package middleware

import (
	"dockernas/internal/config"

	"github.com/gin-gonic/gin"
)

func PathCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if path == "/api/login" || path == "/api/basepath" || path == "/api/systemdir" {
			c.Next()
		} else {
			if config.IsBasePathSet() == false {
				c.JSON(556, gin.H{"msg": "base data path is not set"})
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
