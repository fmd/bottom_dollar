package people

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Background struct {
	Name                string
	Color               string
	HasMiddleName       bool
	Religion            *Religion
	NameChoice          string
	AllowLastNameChange bool
}

func (b *Background) ChangeLastName() bool {
	return b.Religion.ShouldChangeLastName() && b.AllowLastNameChange
}

func (b *Background) ChangeFirstName() bool {
	return b.Religion.ShouldChangeName() && b.Religion.ChangesName
}

func (b *Background) ListChoiceForNameType(t string) string {
	if t == "last" {
		if b.ChangeLastName() {
			return b.Religion.Name
		}
	} else {
		if b.ChangeFirstName() {
			return b.Religion.Name
		}
	}

	return b.NameChoice
}

func NewBackgroundFromChoice(choice BackgroundChoice) *Background {
	b := &Background{}

    religionChoice, err := ReligionChoiceFromList(choice.Religions)
    if err != nil {
        panic(err)
    }

    b.Religion = NewReligionFromChoice(religionChoice)

	b.Name = choice.Name
	b.AllowLastNameChange = choice.AllowLastNameChange
	b.Color = choice.ColorRange[0] //TODO

	index := rand.Intn(len(choice.NameChoices))
	b.NameChoice = choice.NameChoices[index]

	b.HasMiddleName = choice.HasMiddleName
	return b
}
