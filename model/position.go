package model

type Position struct {
	Type  TypeDrink `json:"type"`
	Count int       `json:"count"`
}

type ListPositon struct {
	List []Position `json:"list"`
}
