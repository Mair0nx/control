package control

import (
	"github.com/MSHE97/gitproject2/pkg/types"
)

func Control(card types.Card) {
	if !card.Activity {
		card.Activity = true
	} else {
		card.Activity = false
	}

}
