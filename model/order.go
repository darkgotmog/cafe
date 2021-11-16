package model

type Order struct {
	Id    int        `json:"id"`
	Ready bool       `json:"ready"`
	List  []Position `json:"list"`
}

func NewOrder(id int, flagReady bool, list []Position) *Order {
	return &Order{
		Id:    id,
		Ready: flagReady,
		List:  list,
	}
}
