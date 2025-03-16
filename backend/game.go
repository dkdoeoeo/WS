package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket 升級器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許來自所有來源的連接
	},
}

// 處理 WebSocket 連接的函式
func handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		fmt.Printf("Received: %s\n", p)

		err = conn.WriteMessage(messageType, []byte("Hello from Go WebSocket"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func main() {
	r := gin.Default()

	r.Static("/images", "./image")

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
			return
		}

		handleConnection(conn)
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
