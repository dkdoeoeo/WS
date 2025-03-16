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

	// 不斷讀取客戶端發送的消息
	for {
		// 設定消息的格式
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// 打印收到的消息
		fmt.Printf("Received: %s\n", p)

		// 向客戶端回傳消息
		err = conn.WriteMessage(messageType, []byte("Hello from Go WebSocket"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func main() {
	// 初始化 Gin 引擎
	r := gin.Default()

	// 路由來處理 WebSocket 連接
	r.GET("/ws", func(c *gin.Context) {
		// 使用 Gin 提供的方式獲取請求並升級為 WebSocket 連接
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
			return
		}

		// 處理 WebSocket 連線
		handleConnection(conn)
	})

	// 啟動伺服器
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
