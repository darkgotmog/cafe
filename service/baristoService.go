package service

import (
	"cafe/internal/baristo"
	"cafe/model"
)

func baristoService(chanCashierToBaristo chan model.Order, chanBaristoToCashier chan model.Order) {

	baristo := baristo.NewBaristo(baristo.CokingTimeDefault)

	for v := range chanCashierToBaristo {
		for _, v := range v.List {
			baristo.CokingDrink(v.Type, v.Count)
		}

		v.Ready = true

		chanBaristoToCashier <- v
	}

}

func BaristoService(chanCashierToBaristo chan model.Order, chanBaristoToCashier chan model.Order) {

	go baristoService(chanCashierToBaristo, chanBaristoToCashier)

}
