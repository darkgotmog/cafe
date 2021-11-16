package controller

import (
	"cafe/internal"
	"cafe/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequestOrder godoc
// @Summary post request example
// @Description  request order
// @Accept json
// @Produce plain
// @Param id body  model.OrderId true "OrderId"
// @Success 200 {string} {object} model.Order
// @Router /api/v1/orderReceve [post]

func (cr *Controller) RequestOrder(c echo.Context) error {

	orderId := new(model.OrderId)
	if err := c.Bind(orderId); err != nil {
		return err
	}
	cr.ChanRequestGetOrder <- orderId.Id
	order := <-cr.ChanResponseOrder
	if order.Id != orderId.Id {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, order)
}

// RequestOrderReady godoc
// @Summary post  request orderReady
// @Description  request orderReady
// @Accept json
// @Produce plain
// @Success 200 {string} {object} []model.Order
// @Router /api/v1/orderReady [get]

func (cr *Controller) RequestOrderReady(c echo.Context) error {

	cr.ChanRequest <- internal.ORDERREADY
	oreders := <-cr.ChanResponseOrderReady
	return c.JSON(http.StatusOK, oreders)
}

// RequestOrderWork godoc
// @Summary post  request orderWork
// @Description  request orderWork
// @Accept json
// @Produce plain
// @Success 200 {string} {object} []model.Order
// @Router /api/v1/orderWork [get]

func (cr *Controller) RequestOrderWork(c echo.Context) error {

	cr.ChanRequest <- internal.ORDERWORK
	orders := <-cr.ChanResponseOrderWork
	return c.JSON(http.StatusOK, orders)
}

// RequestNewOrder godoc
// @Summary post request example
// @Description  request order
// @Accept json
// @Produce plain
// @Param message body model.ListPositon true "Add new Order"
// @Success 200 {string} {object} model.OrderId
// @Router /api/v1/orderReceve [post]

func (cr *Controller) RequestNewOrder(c echo.Context) error {

	o := new(model.ListPositon)
	if err := c.Bind(o); err != nil {
		return err
	}
	cr.ChanRequestNewOrder <- *o
	idOrder := <-cr.ChanResponseIdNewOrder

	return c.JSON(http.StatusOK, idOrder)
}
