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
	Name string
	Type TypeDrink
}

func NewDrink(name string, typeDrink TypeDrink) *Drink {
	return &Drink{
		Name: name,
		Type: typeDrink,
	}
}
