package people

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Detective struct {
	SalaryStore int
}

func NewDetective() *Detective {
	d := &Detective{}
	d.SalaryStore = rand.Intn(10000) + 40000
	return d
}

func (d *Detective) String() string {
	return "detective"
}

func (b *Detective) Key() string {
	return "detective"
}

func (d *Detective) Salary() int {
	return d.SalaryStore
}

func (d *Detective) Class() SocioEconomicClass {
	return MIDDLE
}
