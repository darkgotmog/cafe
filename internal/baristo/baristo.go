package baristo

import (
	"cafe/model"
	"fmt"
	"time"
)

var CokingTimeDefault map[model.TypeDrink]int = map[model.TypeDrink]int{
	model.AMERICANO: 100,
	model.CAPUCCINO: 120,
	model.EXPRESSO:  80,
	model.RYSTRETTO: 180,
	model.WATER:     40,
}

type Baristo struct {
	cokingTime map[model.TypeDrink]int
}

func NewBaristo(cokingTime map[model.TypeDrink]int) *Baristo {
	return &Baristo{
		cokingTime: cokingTime,
	}

}
func (b *Baristo) CokingDrink(typeDrink model.TypeDrink, count int) bool {

	timeCoking, ok := b.cokingTime[typeDrink]
	if !ok {
		return false
	}

	time.Sleep(time.Duration(timeCoking*count) * time.Microsecond)
	text := fmt.Sprintf("cokingDrink %+v count: %v", typeDrink, count)
	fmt.Println(text)
	return true
}
