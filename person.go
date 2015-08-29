package main

import (
    "time"
)

type Person struct {
    World *World
    DateOfBirth time.Time
}

func NewPerson(w *World) *Person {
    p := &Person{}
    p.World = w
    p.Birth()
    return p
}

func (p *Person) Birth() {
    p.DateOfBirth = p.World.Date
}