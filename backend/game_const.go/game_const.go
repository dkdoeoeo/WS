package game_const

import (
	"github.com/gorilla/websocket"
	"github.com/timy/WS/backend/card"
)

type Player_state struct {
	Deck_area   []card.Card
	Level_area  []card.Card
	Damage_area []card.Card
	Energy_area []card.Card
	Climax_area []card.Card
	Stage_area  []card.Card
	Rest_area   []card.Card
	Recall_area []card.Card
	Hand_area   []card.Card
	Decks       []card.Deck
}

type GameState struct {
	two_player_states []Player_state
}

type Player struct {
	ID     int             // 玩家 ID
	Conn   *websocket.Conn // WebSocket 連線
	Deckid []int           // 玩家卡組 (卡片 ID 陣列)
}

// 定義接收卡牌選擇的結構
type SaveDeckSelection struct {
	SelectedCardIds []int `json:"selectedCardIds"` // 用來接收卡牌ID的數組
	PlayerID        int   `json:"player_id"`
}
