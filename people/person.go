package main

import (
    "fmt"
)

type Person struct {
    Name *Name
    Gender Gender
    Background *Background
}

func NewPerson() *Person {
    c, err := BackgroundChoiceFromKey("european")
    if err != nil {
        panic(err)
    }

    p := &Person{}
    p.Gender = MALE
    p.Background = NewBackgroundFromChoice(c)

    return p
}

func (p *Person) Describe() {
}