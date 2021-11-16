package cashier

import (
	"cafe/model"
	"errors"
)

var MenuDefault []model.Drink = []model.Drink{
	*model.NewDrink("Вода", model.WATER),
	*model.NewDrink("Капучино", model.CAPUCCINO),
	*model.NewDrink("Эспрессо", model.EXPRESSO),
	*model.NewDrink("Американо", model.AMERICANO),
	*model.NewDrink("Ристретто", model.RYSTRETTO),
}

type Cashier struct {
	Orders         map[int]model.Order
	ListOrderWork  map[int]bool
	ListOrderReady map[int]bool

	lasOrderId int
	listMenu   model.Menu
}

func NewCashier(menu model.Menu) *Cashier {
	return &Cashier{
		Orders:         map[int]model.Order{},
		ListOrderWork:  map[int]bool{},
		ListOrderReady: map[int]bool{},
		lasOrderId:     0,
		listMenu:       menu,
	}
}

func (c *Cashier) Menu() model.Menu {

	return c.listMenu
}

func (c *Cashier) AddOrder(list []model.Position) (int, error) {

	c.lasOrderId += 1

	order := model.NewOrder(c.lasOrderId, false, list)

	c.Orders[order.Id] = *order
	c.ListOrderWork[order.Id] = true

	return order.Id, nil
}

func (c *Cashier) ChangeOrder(order model.Order) error {

	if exists := c.ListOrderWork[order.Id]; exists {
		delete(c.ListOrderWork, order.Id)
	} else {
		return errors.New("Not found Order in ListOrderWork")
	}

	c.Orders[order.Id] = order
	c.ListOrderReady[order.Id] = true

	return nil
}

func (c *Cashier) GetOrder(id int) (model.Order, error) {

	order, exists := c.Orders[id]

	if !exists {
		return order, errors.New("Not found Order")
	}

	if order.Ready {
		delete(c.Orders, id)
		delete(c.ListOrderReady, id)
	}

	return order, nil

}

func (c *Cashier) Order(id int) (model.Order, error) {

	order, exists := c.Orders[id]

	if !exists {
		return order, errors.New("Not found Order")
	}

	return order, nil
}

func (c *Cashier) GetListWorkOrders() []model.Order {

	list := []model.Order{}

	for k := range c.ListOrderWork {

		order, ok := c.Orders[k]
		if ok {
			list = append(list, order)
		}
	}

	return list

}

// RequestMenu godoc
// @Summary Get menu
// @Description Get details of all orders
// @Tags menu
// @Accept  json
// @Produce  json
// @Success 200 {array} nternal.MenuList
// @Router /menu [get]

func (c *Cashier) GetLisReadyOrders() []model.Order {

	list := []model.Order{}

	for k := range c.ListOrderReady {

		order, ok := c.Orders[k]
		if ok {
			list = append(list, order)
		}
	}

	return list

}
