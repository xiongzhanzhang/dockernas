package service

import (
	"io"
	"log"
	"tinycloud/internal/backend/docker"

	"github.com/gorilla/websocket"
)

func ProcessWebsocketConn(conn *websocket.Conn, containerId string, columns string) {
	hr := docker.Exec(containerId, columns)
	defer hr.Close()

	// 退出进程
	defer func() {
		hr.Conn.Write([]byte("exit\r"))
	}()

	log.Println("websocket attach " + containerId)

	// 转发输入/输出至websocket
	go func() {
		wsWriterCopy(hr.Conn, conn)
	}()
	wsReaderCopy(conn, hr.Conn)

	log.Println("websocket disattach " + containerId)
}

// 将终端的输出转发到前端
func wsWriterCopy(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

// 将前端的输入转发到终端
func wsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			writer.Write(p)
		}
	}
}
