package internal

type Position struct {
	Type  TypeDrink `json:"type"`
	Count int       `json:"count"`
}

type Order struct {
	Id    int        `json:"id"`
	Ready bool       `json:"ready"`
	List  []Position `json:"list"`
}

type ListPositon struct {
	List []Position `json:"list"`
}

type OrderId struct {
	Id int `json:"id"`
}

func NewOrder(id int, flagReady bool, list []Position) *Order {
	return &Order{
		Id:    id,
		Ready: flagReady,
		List:  list,
	}
}

type MenuList []Drink
