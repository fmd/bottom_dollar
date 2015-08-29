package main

import (
    "fmt"
)

type Person struct {
    city *City
}

func NewPerson() *Person {
    p := &Person{}
    p.Birth()
    return p
}

func (p *Person) Birth() {
    fmt.Println("Person is born.")
}