package main

import (
    "fmt"
    "time"
)

var world *World

func main() {
    world = NewWorld(time.Date(1870, time.January, 1, 0, 0, 0, 0, time.UTC))
    world.AgeByDaysUntil(time.Date(1971, time.January, 1, 0, 0, 0, 0, time.UTC))

    wait()
}

func wait() {
    var i int
    _, err := fmt.Scanf("%d", &i)

    if err != nil {
        panic(err)
    }
}