package main

import (
	// "errors"
	// "fmt"
	"cafe/internal"
	"cafe/internal/baristo"
	"cafe/internal/cashier"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Cashier *cashier.Cashier
}

func (c *CustomContext) GetCashier() *cashier.Cashier {
	return c.Cashier
}

const (
	MENU int = iota
	ORDERWORK
	ORDERREADY
)

var (
	chanRequest         chan int                  = make(chan int)
	chanRequestGetOrder chan int                  = make(chan int)
	chanRequestOrder    chan int                  = make(chan int)
	chanRequestNewOrder chan internal.ListPositon = make(chan internal.ListPositon)

	chanResponseMenu       chan internal.MenuList = make(chan internal.MenuList)
	chanResponseOrderWork  chan []internal.Order  = make(chan []internal.Order)
	chanResponseOrderReady chan []internal.Order  = make(chan []internal.Order)
	chanResponseOrder      chan internal.Order    = make(chan internal.Order)
	chanResponseIdNewOrder chan internal.OrderId  = make(chan internal.OrderId)
)

func main() {

	chanCashierToBaristo := make(chan internal.Order, 100)
	chanBaristoToCashier := make(chan internal.Order)

	go WorkingBaristo(chanCashierToBaristo, chanBaristoToCashier)
	go WorkingCashier(chanCashierToBaristo, chanBaristoToCashier)

	e := echo.New()

	e.GET("/menu", requestMenu)
	e.GET("/orderWork", requestOrderWork)
	e.GET("/orderReady", requestOrderReady)
	e.POST("/order", requestNewOrder)
	e.POST("/orderReceve", requestOrder)

	e.Logger.Fatal(e.Start(":1323"))

}

func WorkingBaristo(chanCashierToBaristo chan internal.Order, chanBaristoToCashier chan internal.Order) {

	baristo := baristo.NewBaristo(baristo.CokingTimeDefault)

	for v := range chanCashierToBaristo {
		for _, v := range v.List {
			baristo.CokingDrink(v.Type, v.Count)
		}

		v.Ready = true

		chanBaristoToCashier <- v
	}

}

func WorkingCashier(chanCashierToBaristo chan internal.Order, chanBaristoToCashier chan internal.Order) {

	cashier := cashier.NewCashier(cashier.MenuDefault)
	for {
		select {
		case order := <-chanBaristoToCashier:
			{
				cashier.ChangeOrder(order)
			}
		case typeRequest := <-chanRequest:
			{
				switch typeRequest {
				case MENU:
					{
						menu := cashier.Menu()
						chanResponseMenu <- menu
					}
				case ORDERWORK:
					{
						orders := cashier.GetListWorkOrders()
						chanResponseOrderWork <- orders
					}
				case ORDERREADY:
					{
						orders := cashier.GetLisReadyOrders()
						chanResponseOrderReady <- orders
					}

				}

			}
		case listPositon := <-chanRequestNewOrder:
			{

				id, err := cashier.AddOrder(listPositon.List)
				if err == nil {
					chanResponseIdNewOrder <- internal.OrderId{Id: id}

					order, err := cashier.Order(id)
					if err == nil {
						chanCashierToBaristo <- order
					}
				} else {
					chanResponseIdNewOrder <- internal.OrderId{Id: 0}
				}
			}
		case orderId := <-chanRequestGetOrder:
			{
				order, err := cashier.GetOrder(orderId)
				if err == nil {
					chanResponseOrder <- order
				} else {
					chanResponseOrder <- internal.Order{}
				}

			}

		}
	}

}

func requestMenu(c echo.Context) error {

	chanRequest <- MENU
	menu := <-chanResponseMenu
	fmt.Println(menu)

	return c.JSON(http.StatusOK, menu)
}

func requestOrder(c echo.Context) error {

	orderId := new(internal.OrderId)
	if err := c.Bind(orderId); err != nil {
		return err
	}

	chanRequestGetOrder <- orderId.Id
	order := <-chanResponseOrder
	if order.Id != orderId.Id {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, order)
}

func requestOrderReady(c echo.Context) error {

	chanRequest <- ORDERREADY
	oreders := <-chanResponseOrderReady
	return c.JSON(http.StatusOK, oreders)
}

func requestOrderWork(c echo.Context) error {

	chanRequest <- ORDERWORK
	orders := <-chanResponseOrderWork
	return c.JSON(http.StatusOK, orders)
}

func requestNewOrder(c echo.Context) error {

	o := new(internal.ListPositon)
	if err := c.Bind(o); err != nil {
		return err
	}
	chanRequestNewOrder <- *o
	idOrder := <-chanResponseIdNewOrder

	return c.JSON(http.StatusOK, idOrder)
}
