package baristo

import (
	"cafe/internal"
	"fmt"
	"time"
)

var CokingTimeDefault map[internal.TypeDrink]int = map[internal.TypeDrink]int{
	internal.AMERICANO: 100,
	internal.CAPUCCINO: 120,
	internal.EXPRESSO:  80,
	internal.RYSTRETTO: 180,
	internal.WATER:     40,
}

type Baristo struct {
	cokingTime map[internal.TypeDrink]int
}

func NewBaristo(cokingTime map[internal.TypeDrink]int) *Baristo {
	return &Baristo{
		cokingTime: cokingTime,
	}

}
func (b *Baristo) CokingDrink(typeDrink internal.TypeDrink, count int) bool {

	timeCoking, ok := b.cokingTime[typeDrink]
	if !ok {
		return false
	}

	time.Sleep(time.Duration(timeCoking*count) * time.Microsecond)
	text := fmt.Sprintf("cokingDrink %+v count: %v", typeDrink, count)
	fmt.Println(text)
	return true
}
