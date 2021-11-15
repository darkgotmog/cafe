package main

import (
	// "errors"
	// "fmt"
	config "cafe/configs"
	"cafe/internal"
	"cafe/internal/baristo"
	"cafe/internal/cashier"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	chanRequestNewOrder chan internal.ListPositon = make(chan internal.ListPositon)
	chanRequestGetOrder chan int                  = make(chan int)

	chanResponseMenu       chan internal.MenuList = make(chan internal.MenuList)
	chanResponseOrderWork  chan []internal.Order  = make(chan []internal.Order)
	chanResponseOrderReady chan []internal.Order  = make(chan []internal.Order)
	chanResponseOrder      chan internal.Order    = make(chan internal.Order)
	chanResponseIdNewOrder chan internal.OrderId  = make(chan internal.OrderId)
)

const pathConfig string = "config.ini"

// @title API
// @version 1.0
// @description Swagger API for Golang.
// @host localhost:1323
// @BasePath /api/v1

func main() {

	conf, err := config.LoadConfig(pathConfig)
	if err != nil {
		os.Exit(1)
	}

	chanCashierToBaristo := make(chan internal.Order, 100)
	chanBaristoToCashier := make(chan internal.Order)

	go WorkingBaristo(chanCashierToBaristo, chanBaristoToCashier)
	go WorkingCashier(chanCashierToBaristo, chanBaristoToCashier)

	StartServer(conf.Port)

}

func StartServer(port int) {
	e := echo.New()

	v1 := e.Group("/api/v1")
	{
		v1.GET("/menu", requestMenu)
		v1.GET("/orderWork", requestOrderWork)
		v1.GET("/orderReady", requestOrderReady)
		v1.POST("/order", requestNewOrder)
		v1.POST("/orderReceve", requestOrder)
	}

	address := fmt.Sprintf(":%v", port)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(address))

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

// requestMenu godoc
// @Summary Get menu
// @Description Get details of all orders
// @Tags menu
// @Accept  json
// @Produce  json
// @Success 200 {array} nternal.MenuList
// @Router /menu [get]

func requestMenu(c echo.Context) error {

	chanRequest <- MENU
	menu := <-chanResponseMenu

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
