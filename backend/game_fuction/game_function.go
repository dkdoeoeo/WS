package game_function

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/timy/WS/backend/game_const.go"
)

func UpdateDeck(c *gin.Context) {
	var deckSelection game_const.DeckSelection

	if err := c.ShouldBindJSON(&deckSelection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的請求資料"})
		return
	}

	fmt.Printf("收到選擇的卡牌 ID: %v\n", deckSelection.SelectedCardIds)

	c.JSON(http.StatusOK, gin.H{"message": "卡牌已成功提交"})

}

func Init_Player_state() game_const.Player_state {
	new_Player_state := game_const.Player_state{}
	return new_Player_state
}
