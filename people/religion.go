package people

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Religion struct {
	Name                       string
	Plural                     string
	Singular                   string
	ChangesName                bool
	ChangesLastName            bool
	ChangesNamesOnNonReligious bool
	Religious                  bool
}

func (r *Religion) Status() string {
	if r.Religious {
		return "religious"
	}
	return "non-religious"
}

func (r *Religion) ShouldChangeLastName() bool {
	return (r.Religious || r.ChangesNamesOnNonReligious) && r.ChangesLastName
}

func (r *Religion) ShouldChangeName() bool {
	return (r.Religious || r.ChangesNamesOnNonReligious) && r.ChangesName
}

func NewReligionFromChoice(choice ReligionChoice) *Religion {
	r := &Religion{}
	r.Name = choice.Name
	r.Plural = choice.Plural
	r.Singular = choice.Singular
	r.ChangesName = choice.ChangesName
	r.ChangesLastName = choice.ChangesLastName
	r.ChangesNamesOnNonReligious = choice.ChangesNamesOnNonReligious

	chance := rand.Intn(100)
	if chance < choice.ReligiousPercent {
		r.Religious = true
	}

	return r
}
