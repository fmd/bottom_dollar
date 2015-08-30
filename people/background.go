package people

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Background struct {
	Name          string
	Color         string
	Immigrated    bool
	Naturalized   bool
	HasMiddleName bool
	Religion      *Religion
}

func NewBackgroundFromChoice(choice BackgroundChoice) *Background {
	b := &Background{}
	b.Name = choice.Name
	b.Color = choice.ColorRange[0] //TODO

	chance := rand.Intn(100)
	if chance < choice.ImmigrationPercent {
		b.Immigrated = true
	}

	chance = rand.Intn(100)
	if chance < choice.NaturalizationPercent {
		b.Naturalized = true
	}

	b.HasMiddleName = choice.HasMiddleName
	religionChoice, err := ReligionChoiceFromList(choice.Religions)
	if err != nil {
		panic(err)
	}

	b.Religion = NewReligionFromChoice(religionChoice)

	return b
}
