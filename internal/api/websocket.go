package api

import (
	"log"
	"net/http"
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func InstanceWebTerminal(c *gin.Context) {
	if !service.IsTokenValid(c.Query("token")) {
		c.JSON(555, gin.H{"msg": "Authentication error"})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	service.ProcessWebsocketConn(ws, c.Query("containerid"), c.Query("columns"))
}
