package main

import (
    "fmt"
    "time"
)

var world *World
var city *City

func main() {
    world = NewWorld(time.Date(1870, time.January, 1, 0, 0, 0, 0, time.UTC))
    city = NewCity()

    fmt.Println(world, city)
    wait()
}

func wait() {
    var i int
    _, err := fmt.Scanf("%d", &i)

    if err != nil {
        panic(err)
    }
}