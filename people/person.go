package people

import (
    "fmt"
)

type Person struct {
	Name       *Name
	Gender     Gender
	Background *Background
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
    fmt.Println(p.Name.Full())
}
