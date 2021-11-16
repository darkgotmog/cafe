package controller

import (
	"cafe/internal"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequestMenu godoc
// @Summary post request example
// @Description  request menu
// @Accept json
// @Produce plain
// @Param message body model.Menu true "Menu list"
// @Success 200 {string}
// @Router /api/v1/menu [get]

func (cr *Controller) RequestMenu(c echo.Context) error {

	cr.ChanRequest <- internal.MENU
	menu := <-cr.ChanResponseMenu

	return c.JSON(http.StatusOK, menu)
}
