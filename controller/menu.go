package controller

import (
	"cafe/internal"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequestMenu godoc
// @Summary RequestMenu
// @Description  request menu
// @Accept plain
// @Produce json
// @Success 200 {object} model.Menu
// @Router /menu [get]

func (cr *Controller) RequestMenu(c echo.Context) error {

	cr.ChanRequest <- internal.MENU
	menu := <-cr.ChanResponseMenu

	return c.JSON(http.StatusOK, menu)
}
