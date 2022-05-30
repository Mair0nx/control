package main

import (
	"fmt"

	"github.com/MSHE97/gitproject2/pkg/types"
	"github.com/Mair0nx/control/pkg/control"
)

func main() {

	var card types.Card

	card.Activity = true

	control.Control(card)

	fmt.Printf("%+v", card)
}
