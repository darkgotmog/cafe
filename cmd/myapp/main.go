package main

import (

	// "fmt"

	config "cafe/configs"
	"cafe/controller"
	"cafe/model"
	"cafe/service"
	"os"

	"fmt"

	_ "cafe/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const pathConfig string = "config.ini"

func main() {

	conf, err := config.LoadConfig(pathConfig)
	if err != nil {
		os.Exit(1)
	}

	cr := controller.NewController()

	chanCashierToBaristo := make(chan model.Order, 100)
	chanBaristoToCashier := make(chan model.Order)

	service.BaristoService(chanCashierToBaristo, chanBaristoToCashier)
	service.CashierService(chanCashierToBaristo, chanBaristoToCashier, cr)

	StartServer(conf.Port, cr)

}

// @title API
// @version 1.0
// @description Cafe API .
// @host localhost:1323
// @BasePath /api/v1

func StartServer(port int, cr *controller.Controller) {
	e := echo.New()

	v1 := e.Group("/api/v1")
	{
		v1.GET("/menu", cr.RequestMenu)
		v1.GET("/orderWork", cr.RequestOrderWork)
		v1.GET("/orderReady", cr.RequestOrderReady)
		v1.POST("/order", cr.RequestNewOrder)
		v1.POST("/orderReceve", cr.RequestOrder)
	}

	address := fmt.Sprintf(":%v", port)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(address))

}
