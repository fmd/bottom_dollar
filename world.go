package main

import (
    "fmt"
    "time"
    "math/rand"
)

const STARTING_POPULATION int = 30000
const MAX_RANDOM_STARTING_POPULATION int = 5000

func init() {
    rand.Seed(time.Now().UnixNano())
}

type World struct {
    Date time.Time
    Ticker *time.Ticker
    Paused bool
}

func NewWorld(date time.Time) *World {
    w := &World{}
    w.Date = date
    w.Ticker = time.NewTicker(time.Second * 1)
    w.Paused = true
    go w.advanceTicker()
    return w
}

func (w *World) Generate() {
    w.generateInitialPopulation()
    w.ageByDaysUntil(time.Date(1971, time.January, 1, 0, 0, 0, 0, time.UTC))
}

func (w *World) advanceTicker() {
    for _ = range w.Ticker.C {
        if !w.Paused {
            w.tick(0, 0, 30)
        }
    }
}

func (w *World) generateInitialPopulation() {
    population := STARTING_POPULATION
    plusRandom := rand.Intn(MAX_RANDOM_STARTING_POPULATION)

    var person *Person
    for i := 0; i < population + plusRandom; i++ {
        person = NewPerson()
        person.RandomizeAge()
    }

    fmt.Println(population + plusRandom, "people live in the city.")
}

func (w *World) age(years int, months int, days int) {
    w.Date = w.Date.AddDate(years, months, days)
}

func (w *World) tick(hours time.Duration, minutes time.Duration, seconds time.Duration) {
    w.Date = w.Date.Add(time.Second * seconds + time.Minute * minutes + time.Hour * hours)
}

func (w *World) ageByDaysUntil(date time.Time) {
    for w.Date.Before(date) {
        w.age(0, 0, 1)
    }
}