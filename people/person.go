package people

import (
    "fmt"
    "github.com/fmd/bottom_dollar/professions"
)

type Person struct {
	Name       *Name
	Gender     Gender
	Background *Background
    Profession Profession
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
    p.Gender = RandomGender()
    p.Background = b
    p.Name = NameFromBackgroundAndGender(p.Gender, p.Background)
    return p
}

func (p *Person) Describe() {
    s := p.Name.Full() + " -> " + p.Gender.String()
    s = s + ", a " + p.Background.Religion.Status() + " " + p.Background.Class.String() + " "
    s = s + p.Background.Name + " " + p.Background.Religion.Singular + "."
    fmt.Println(s)
}
