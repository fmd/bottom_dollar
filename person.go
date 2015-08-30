package main

import (
    "fmt"
)

type Person struct {
    Name *Name
    Gender string
    Background *Background
}

func NewPerson() *Person {
    return &Person{}
}

func (p *Person) Say(message string) {
    fmt.Println(message)
}