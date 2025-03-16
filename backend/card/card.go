package card

type Card struct {
	Name           string         // 名稱
	Level          int            // 等級
	Cost           int            // 代價
	Counterattack  bool           // 是否可以反擊
	Special_effect Special_effect // 特殊符號
	Card_effect    Card_effect    //卡牌效果
	Power          int            //戰鬥力
	Damage_points  int            //魂傷
	Feature        []string       //特徵
	Color          int            //1紅色、2藍色、3黃色
	ImagePath      string
}

type Deck struct {
	Cards []Card // 卡片陣列
}

func NewCard(name string, level int, cost int, counterattack bool, spcial_effect Special_effect, card_effect Card_effect, power int, damage_points int, feature []string, color int, imagePath string) *Card {
	return &Card{
		Name:           name,
		Level:          level,
		Cost:           cost,
		Counterattack:  counterattack,
		Special_effect: spcial_effect,
		Card_effect:    card_effect,
		Power:          power,
		Damage_points:  damage_points,
		Feature:        feature,
		Color:          color,
		ImagePath:      imagePath,
	}
}

func LoadCards() *Deck {
	// 創建一個牌庫
	deck := &Deck{}

	return deck
}

// 為牌庫新增卡片
func (d *Deck) AddCard(card *Card) {
	d.Cards = append(d.Cards, *card)
}
