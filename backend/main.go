package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/timy/WS/backend/game_const.go"
)

// WebSocket 升級器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許來自所有來源的連接
	},
}

var (
	rooms          = make(map[string]game_const.Player_state) // 房間狀態
	onlinePlayers  = make(map[int]*game_const.Player)         // 線上玩家隊列
	waitingPlayers = []int{}
	MatchQueue     []int
	mu             sync.Mutex // 保護共享資源
	roomIDCounter  = 1        // 房間編號
	curPlayerId    = 0        //當前空閒玩家id
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
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// 解析 JSON 訊息
		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}

		switch msg["type"].(string) {
		case "player_id":
			playerId, err := strconv.Atoi(fmt.Sprintf("%v", msg["player_id"]))
			if err != nil || playerId == -1 {
				fmt.Println("無效的玩家 ID:", msg["player_id"])
				return
			}

			mu.Lock()
			onlinePlayers[playerId].Conn = conn
			waitingPlayers = append(waitingPlayers, playerId)
			mu.Unlock()

			fmt.Println("收到玩家 ID:", playerId)
			fmt.Println("目前等待配對的玩家:", waitingPlayers)

			if len(waitingPlayers) >= 2 {
				player1conn := onlinePlayers[waitingPlayers[0]].Conn
				player2conn := onlinePlayers[waitingPlayers[1]].Conn

				roomID := createRoom()
				rooms[roomID] = Init_Player_state()

				player1conn.WriteJSON(gin.H{"playerID": 0})
				player2conn.WriteJSON(gin.H{"playerID": 1})

				// 發送初始遊戲狀態給兩位玩家
				player1conn.WriteJSON(rooms[roomID])
				player2conn.WriteJSON(rooms[roomID])

				waitingPlayers = waitingPlayers[2:]

				go listenForPlayerInput(player1conn, roomID)
				go listenForPlayerInput(player2conn, roomID)
				go updateGameLoop(player1conn, player2conn, roomID)
			}

		default:
			fmt.Println("未知訊息類型:", msg["type"])
		}
	}
}

func listenForPlayerInput(conn *websocket.Conn, roomID string) {
	defer conn.Close()
}

func updateGameLoop(Player1_conn *websocket.Conn, Player2_conn *websocket.Conn, roomID string) {
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

	r.POST("/updateDeck", Save_Player_Deck)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func Save_Player_Deck(c *gin.Context) {
	var deckSelection game_const.SaveDeckSelection

	if err := c.ShouldBindJSON(&deckSelection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的請求資料"})
		return
	}

	fmt.Printf("玩家 ID: %v\n", deckSelection.PlayerID)
	fmt.Printf("收到選擇的卡牌 ID: %v\n", deckSelection.SelectedCardIds)

	if player, exists := onlinePlayers[deckSelection.PlayerID]; exists {
		fmt.Printf("玩家 ID存在%v\n", deckSelection.PlayerID)
		player.Deckid = deckSelection.SelectedCardIds // 存起來，等待匹配
	} else {
		fmt.Printf("玩家 ID不存在，發送ID: %d\n", curPlayerId)
		onlinePlayers[curPlayerId] = &game_const.Player{
			ID:     curPlayerId,
			Deckid: deckSelection.SelectedCardIds,
		}
		c.JSON(http.StatusOK, gin.H{"curPlayerId": curPlayerId})
		curPlayerId++
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "卡牌已成功提交"})
}

func Init_Player_state() game_const.Player_state {
	new_Player_state := game_const.Player_state{}
	return new_Player_state
}
