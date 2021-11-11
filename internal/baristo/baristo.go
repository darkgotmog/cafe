package baristo

import (
	"cafe/internal"
	"fmt"
	"time"
)

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

	time.Sleep(time.Duration(timeCoking) * time.Millisecond)
	fmt.Printf("cokingDrink %v count: %v", typeDrink, count)
	return true
}

