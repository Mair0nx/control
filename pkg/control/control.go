package control

import ("fmt"
"github.com/MSHE97/gitproject2/pkg/types")

var card types.Card

func control(card types.Card) {
if !card.Activity {
	card.Activity = true
} else {
	card.Activity = false
}
fmt.Println(card)
}