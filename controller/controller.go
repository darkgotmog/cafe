package controller

import (
	"cafe/model"
)

// Controller example
type Controller struct {
	ChanRequest         chan int
	ChanRequestNewOrder chan model.ListPositon
	ChanRequestGetOrder chan int

	ChanResponseMenu       chan model.Menu
	ChanResponseOrderWork  chan []model.Order
	ChanResponseOrderReady chan []model.Order
	ChanResponseOrder      chan model.Order
	ChanResponseIdNewOrder chan model.OrderId
}

// NewController example
func NewController() *Controller {
	return &Controller{
		ChanRequest:         make(chan int),
		ChanRequestNewOrder: make(chan model.ListPositon),
		ChanRequestGetOrder: make(chan int),

		ChanResponseMenu:       make(chan model.Menu),
		ChanResponseOrderWork:  make(chan []model.Order),
		ChanResponseOrderReady: make(chan []model.Order),
		ChanResponseOrder:      make(chan model.Order),
		ChanResponseIdNewOrder: make(chan model.OrderId),
	}
}
