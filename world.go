package main

import (
    "fmt"
    "time"
)

type World struct {
    Date time.Time
    Ticker *time.Ticker
    Paused bool
}

func NewWorld(date time.Time) *World {
    w := &World{}
    w.Date = date
    w.Ticker = time.NewTicker(time.Second * 1)
    w.Paused = false

    go func() {
        for _ = range w.Ticker.C {
            if !w.Paused {
                w.Tick(0, 0, 30)
                fmt.Println(w.Date)
            }
        }
    }()

    return w
}

func (w *World) Age(years int, months int, days int) {
    w.Date = w.Date.AddDate(years, months, days)
}

func (w *World) Tick(hours time.Duration, minutes time.Duration, seconds time.Duration) {
    w.Date = w.Date.Add(time.Second * seconds + time.Minute * minutes + time.Hour * hours)
}