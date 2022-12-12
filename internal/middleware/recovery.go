package middleware

import "github.com/gin-gonic/gin"

func defaultHandleRecovery(c *gin.Context, err any) {
	if e, ok := err.(error); ok {
		c.JSON(500, gin.H{"msg": e.Error()})
	} else {
		c.JSON(500, gin.H{"msg": "server error"})
	}

	c.Abort()
}

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(defaultHandleRecovery)
}
