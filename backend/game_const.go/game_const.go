package game_const

import "github.com/timy/WS/backend/card"

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
	Deck        card.Deck
}

type GameState struct {
	players []Player_state
}

// 定義接收卡牌選擇的結構
type DeckSelection struct {
	SelectedCardIds []int `json:"selectedCardIds"` // 用來接收卡牌ID的數組
}
