package people

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Religion struct {
	Name      string
	Plural    string
	Religious bool
}

func NewReligionFromChoice(choice ReligionChoice) *Religion {
	r := &Religion{}
	r.Name = choice.Name
	r.Plural = choice.Plural

	chance := rand.Intn(99) + 1
	if chance < choice.ReligiousPercent {
		r.Religious = true
	}

	return r
}
