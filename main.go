package main

import (
    "time"
    "math/rand"
	"github.com/fmd/bottom_dollar/people"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
	p := people.NewPersonFromBackgroundChoice(people.RandomBackgroundChoice())
	p.Describe()
}
