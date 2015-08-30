package main

import (
    "time"
    "math/rand"
)

type SocioEconomicClass int

const (
    HOMELESS SocioEconomicClass = iota
    WORKING
    LOWER_MIDDLE
    MIDDLE
    UPPER_MIDDLE
    UPPER
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type Person struct {
    DateOfBirth time.Time
    Relationships []*Relationship
    Class SocioEconomicClass
}

func NewPerson() *Person {
    p := &Person{}
    p.Birth()
    people = append(people, p)
    return p
}

func (p *Person) Birth() {
    p.DateOfBirth = world.Date
}

func getRandomYearsBack() int {
    norm := rand.NormFloat64() * 15 + 35
    if norm < 18 {
        return getRandomYearsBack()
    }
    return int(norm)
}

func (p *Person) RandomizeAge() {
    p.DateOfBirth = p.DateOfBirth.AddDate(-getRandomYearsBack(), 0, rand.Intn(365))
}