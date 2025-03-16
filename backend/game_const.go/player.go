package game_const

import "github.com/timy/WS/backend/card"

type player struct {
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
