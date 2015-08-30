package main

import (
    "time"
)

var world *World
var people []*Person

func main() {
    world = NewWorld(time.Date(1870, time.January, 1, 0, 0, 0, 0, time.UTC))
    world.Generate()
}