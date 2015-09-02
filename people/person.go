package people

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Person struct {
	Name       *Name
	Gender     Gender
	Background *Background
	Profession Profession
}

func (p *Person) AssignRandomProfession() {
	professions := []Profession{&Bartender{},
								&Detective{},
								&Jobless{}}

	index := rand.Intn(len(professions))
	chosen := professions[index]
	if p.Background.Religion.AllowsProfessionKey(chosen.Key()) {
		p.Profession = chosen
	} else {
		p.AssignRandomProfession()
	}
}

func NewPersonFromBackgroundKey(backgroundKey string) *Person {
	c, err := BackgroundChoiceFromKey(backgroundKey)
	if err != nil {
		panic(err)
	}

	return NewPersonFromBackgroundChoice(c)
}

func NewPersonFromBackgroundChoice(choice BackgroundChoice) *Person {
	b := NewBackgroundFromChoice(choice)
	return NewPersonFromBackground(b)
}

func NewPersonFromBackground(b *Background) *Person {
	p := &Person{}
	p.Background = b
	p.Gender = RandomGender()
	p.Name = NameFromBackgroundAndGender(p.Gender, p.Background)
	p.AssignRandomProfession()
	return p
}

func (p *Person) Describe() {
	s := p.Name.Full() + " -> " + p.Gender.String()
	s = s + ", a " + p.Background.Religion.Status() + " " + p.Profession.Class().String() + " "
	s = s + p.Background.Name + " " + p.Background.Religion.Name + " " + p.Profession.String() + "."
	fmt.Println(s)
}
