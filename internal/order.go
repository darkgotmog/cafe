package internal

type Position struct {
	Type  TypeDrink
	Count int
}

type Order struct {
	Id    int
	Ready bool
	List  []Position
}

func NewOrder(id int, flagReady bool, list []Position) *Order {
	return &Order{
		Id:    id,
		Ready: flagReady,
		List:  list,
	}
}

type MenuList []Drink
