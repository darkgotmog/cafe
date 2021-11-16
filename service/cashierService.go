package service

import (
	"cafe/controller"
	"cafe/internal"
	"cafe/internal/cashier"
	"cafe/model"
)

func CashierService(chanCashierToBaristo chan model.Order, chanBaristoToCashier chan model.Order, cr *controller.Controller) {
	go cashierService(chanCashierToBaristo, chanBaristoToCashier, cr)
}

func cashierService(chanCashierToBaristo chan model.Order, chanBaristoToCashier chan model.Order, cr *controller.Controller) {

	cashier := cashier.NewCashier(cashier.MenuDefault)
	for {
		select {
		case order := <-chanBaristoToCashier:
			{
				cashier.ChangeOrder(order)
			}
		case typeRequest := <-cr.ChanRequest:
			{
				switch typeRequest {
				case internal.MENU:
					{
						menu := cashier.Menu()
						cr.ChanResponseMenu <- menu
					}
				case internal.ORDERWORK:
					{
						orders := cashier.GetListWorkOrders()
						cr.ChanResponseOrderWork <- orders
					}
				case internal.ORDERREADY:
					{
						orders := cashier.GetLisReadyOrders()
						cr.ChanResponseOrderReady <- orders
					}

				}

			}
		case listPositon := <-cr.ChanRequestNewOrder:
			{

				id, err := cashier.AddOrder(listPositon.List)
				if err == nil {
					cr.ChanResponseIdNewOrder <- model.OrderId{Id: id}

					order, err := cashier.Order(id)
					if err == nil {
						chanCashierToBaristo <- order
					}
				} else {
					cr.ChanResponseIdNewOrder <- model.OrderId{Id: 0}
				}
			}
		case orderId := <-cr.ChanRequestGetOrder:
			{
				order, err := cashier.GetOrder(orderId)
				if err == nil {
					cr.ChanResponseOrder <- order
				} else {
					cr.ChanResponseOrder <- model.Order{}
				}

			}

		}
	}

}
