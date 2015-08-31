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
        b, err := people.BackgroundChoiceFromKey("african")
        if err != nil {
            panic(err)
        }
		p := people.NewPersonFromBackgroundChoice(b)
		p.Profession = professions.NewDetective()
		p.Describe()
	}
}