package professions

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
    d := &Bartender{}
    d.SalaryStore = rand.Intn(6000) + 10000
    return d
}

func (d *Bartender) String() string {
    return "bartender"
}

func (d *Bartender) Salary() int {
    return d.SalaryStore
}

func (d *Bartender) Class() SocioEconomicClass {
    return WORKING
}
