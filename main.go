package main

import (
    "github.com/fmd/bottom_dollar/places"
    "fmt"
)

func main() {
    w := NewWindow()
    var shouldQuit bool

    MakeBar()

    for !shouldQuit {
        shouldQuit = ProcessOneFrameOfInput()
        w.ProcessOneFrame()
    }

    w.Destroy()
}

func MakeBar() {
    p := places.Place{}
    p.Layers = append(p.Layers, places.NewExampleGroundLayer())
    p.Layers = append(p.Layers, places.NewExampleWallLayer())

    fmt.Println(p.Layers)
}