package controller

import (
	"cafe/internal"
	"cafe/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequestOrder godoc
// @Summary RequestOrder
// @Description  request order
// @Accept json
// @Produce json
// @Param input body  model.OrderId true "OrderId"
// @Success 200 {object} model.Order
// @Failure 404 {string} string "Not Found"
// @Router /orderReceve [post]

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
// @Summary pRequestOrderReady
// @Description  request orderReady
// @Accept plain
// @Produce json
// @Success 200 {object} []model.Order
// @Router /orderReady [get]

func (cr *Controller) RequestOrderReady(c echo.Context) error {

	cr.ChanRequest <- internal.ORDERREADY
	oreders := <-cr.ChanResponseOrderReady
	return c.JSON(http.StatusOK, oreders)
}

// RequestOrderWork godoc
// @Summary RequestOrderWork
// @Description  request orderWork
// @Accept plain
// @Produce json
// @Success 200 {object} []model.Order
// @Router /orderWork [get]

func (cr *Controller) RequestOrderWork(c echo.Context) error {

	cr.ChanRequest <- internal.ORDERWORK
	orders := <-cr.ChanResponseOrderWork
	return c.JSON(http.StatusOK, orders)
}

// RequestNewOrder godoc
// @Summary RequestNewOrder
// @Description  request order
// @Accept json
// @Produce json
// @Param input body model.ListPositon true "Add new Order"
// @Success 200 {object} model.OrderId
// @Router /orderReceve [post]

func (cr *Controller) RequestNewOrder(c echo.Context) error {

	o := new(model.ListPositon)
	if err := c.Bind(o); err != nil {
		return err
	}
	cr.ChanRequestNewOrder <- *o
	idOrder := <-cr.ChanResponseIdNewOrder

	return c.JSON(http.StatusOK, idOrder)
}
