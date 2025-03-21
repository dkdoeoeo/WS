package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/timy/WS/backend/game_const.go"
	game_function "github.com/timy/WS/backend/game_fuction"
)

// WebSocket 升級器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許來自所有來源的連接
	},
}

var (
	rooms         = make(map[string]game_const.Player_state) // 房間狀態
	players       = make([]*websocket.Conn, 0)               // 等待匹配的玩家隊列
	mu            sync.Mutex                                 // 保護共享資源
	roomIDCounter = 1                                        // 房間編號
	curRoomNum    = 0
)

func createRoom() string {
	mu.Lock()
	defer mu.Unlock()

	// 創建一個新的房間並返回房間ID
	roomID := fmt.Sprintf("room-%d", roomIDCounter)
	roomIDCounter++
	return roomID
}

// 處理 WebSocket 連接的函式
func handleConnection(conn *websocket.Conn) {
	mu.Lock()
	players = append(players, conn)
	mu.Unlock()

	if len(players) >= 2 {
		player1 := players[0]
		player2 := players[1]

		roomID := createRoom()
		curRoomNum++

		rooms[roomID] = game_function.Init_Player_state()

		player1.WriteJSON(gin.H{"playerID": 0})
		player2.WriteJSON(gin.H{"playerID": 1})

		// 配對成功後通知玩家
		player1.WriteJSON(gin.H{"message": "配對成功，前往遊戲！"})
		player2.WriteJSON(gin.H{"message": "配對成功，前往遊戲！"})

		// 發送初始遊戲狀態給兩位玩家
		player1.WriteJSON(rooms[roomID])
		player2.WriteJSON(rooms[roomID])

		players = players[2:]
	} else {
		conn.WriteJSON(map[string]string{"message": "Waiting for another player"})
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

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

	r.POST("/updateDeck", game_function.UpdateDeck)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
