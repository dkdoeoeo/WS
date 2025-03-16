package card

import "fmt"

type Special_effect func(card *Card) string

type Card_effect func(card *Card) string

// AttackEffect 定義攻擊效果
func AttackEffect(card *Card) string {
	return fmt.Sprintf("Card (Level %d) attacks with %d attack power!", card.Level)
}

// HealEffect 定義治療效果
func HealEffect(card *Card) string {
	return fmt.Sprintf("Card (Level %d) heals for %d points!", card.Level)
}
