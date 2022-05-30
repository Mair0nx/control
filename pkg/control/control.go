package control

import (
	"github.com/MSHE97/gitproject2/pkg/types"
)

func Control(card *types.Card) {
	//Функция на проверку активности карты, и замены её на противоположное значение
	if !(*card).Activity {
		(*card).Activity = true
	} else {
		(*card).Activity = false
	}

}
