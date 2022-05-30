package main

import (
	"fmt"

	"github.com/MSHE97/gitproject2/pkg/types"
	"github.com/Mair0nx/control/pkg/control"
)

func main() {

	var card types.Card

	//Карта, статус которой активен || Не активен
	card.Activity = false

	//Вызов конкретной карты в функцию и смена её значения
	control.Control(&card)

	//Вывод форматированного статуска карты
	fmt.Printf("%+v", card)
}
