package professions

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
    d.SalaryStore = rand.Intn(6000) + 10000
    return d
}

func (d *Detective) String() string {
    return "bartender"
}

func (d *Detective) Salary() int {
    return d.SalaryStore
}

func (d *Detective) Class() SocioEconomicClass {
    return WORKING
}
