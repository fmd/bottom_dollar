package main

import (
    "time"
)

type City struct {}

func NewCity() *City {
    c := &City{}
    c.PopulateUntil(time.Date(1971, time.January, 1, 0, 0, 0, 0, time.UTC))
    return c
}

func (c *City) PopulateUntil(date time.Time) {
    for world.Date.Before(date) {
        world.Date = world.Date.AddDate(0, 1, 0)
    }
}