package main

import (
	"github.com/fmd/bottom_dollar/people"
	"github.com/fmd/bottom_dollar/professions"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i := 0; i < 2000; i++ {
		p := people.NewPersonFromBackgroundChoice(people.RandomBackgroundChoice())
		p.Profession = professions.NewDetective()
		p.Describe()
	}
}
