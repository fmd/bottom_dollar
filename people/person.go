package people

import (
    "fmt"
)

type Person struct {
	Name       *Name
	Gender     Gender
	Background *Background
}

func NewPerson() *Person {
	c, err := BackgroundChoiceFromKey("irish")
	if err != nil {
		panic(err)
	}

	p := &Person{}
	p.Gender = FEMALE
	p.Background = NewBackgroundFromChoice(c)
    p.Name = NameFromBackgroundAndGender(p.Gender, p.Background)

	return p
}

func (p *Person) Describe() {
    fmt.Println(p.Name)
}
