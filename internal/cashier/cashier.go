package cashier

import (
	"cafe/internal"
	"errors"
)

type Cashier struct {
	Orders         map[int]internal.Order
	ListOrderWork  map[int]bool
	ListOrderReady map[int]bool

	lasOrderId int
	listMenu   internal.MenuList
}

func NewCashier(listMenu internal.MenuList) *Cashier {
	return &Cashier{
		listMenu: listMenu,
	}
}

func (c *Cashier) Menu(typeDrink internal.TypeDrink, count int) internal.MenuList {

	return c.listMenu
}

func (c *Cashier) AddOrder(list []internal.Position) (int, error) {

	c.lasOrderId += 1

	order := internal.NewOrder(c.lasOrderId, false, list)

	c.Orders[order.Id] = *order
	c.ListOrderWork[order.Id] = true

	return order.Id, nil
}

func (c *Cashier) ChangeOrder(order internal.Order) error {

	if exists := c.ListOrderWork[order.Id]; exists {
		delete(c.ListOrderWork, order.Id)
	} else {
		return errors.New("Not found Order in ListOrderWork")
	}

	c.Orders[order.Id] = order
	c.ListOrderReady[order.Id] = true

	return nil
}

func (c *Cashier) GetOrder(id int) (internal.Order, error) {

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

func (c *Cashier) GetListWorkOrders(id int) []internal.Order {

	list := []internal.Order{}

	for k := range c.ListOrderWork {

		order, ok := c.Orders[k]
		if ok {
			list = append(list, order)
		}
	}

	return list

}

func (c *Cashier) GetLisReadyOrders(id int) []internal.Order {

	list := []internal.Order{}

	for k := range c.ListOrderReady {

		order, ok := c.Orders[k]
		if ok {
			list = append(list, order)
		}
	}

	return list

}
