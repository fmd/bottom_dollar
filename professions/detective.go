package professions

import (
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type Detective struct {
    Salary int
}

func NewDetective() *Detective {
    d := &Detective{}
    d.Salary = rand.Intn(10000) + 40000
    return d
}

func (d *Detective) String() {
    return "detective"
}

func (d *Detective) Salary() {
    return d.Salary
}