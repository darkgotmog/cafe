package internal

type TypeDrink int

const (
	WATER TypeDrink = iota
	CAPUCCINO
	EXPRESSO
	AMERICANO
	RYSTRETTO
)

type Drink struct {
	Name string    `json:"name"`
	Type TypeDrink `json:"type"`
}

func NewDrink(name string, typeDrink TypeDrink) *Drink {
	return &Drink{
		Name: name,
		Type: typeDrink,
	}
}
