package people

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Bartender struct {
	SalaryStore int
}

func NewBartender() *Bartender {
	b := &Bartender{}
	b.SalaryStore = rand.Intn(6000) + 10000
	return b
}

func (b *Bartender) String() string {
	return "bartender"
}

func (b *Bartender) Key() string {
	return "bartender"
}

func (b *Bartender) Salary() int {
	return b.SalaryStore
}

func (b *Bartender) Class() SocioEconomicClass {
	return WORKING
}
