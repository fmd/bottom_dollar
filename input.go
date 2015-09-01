package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
)

var event sdl.Event

func ProcessOneFrameOfInput() bool {
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        switch t := event.(type) {
        case *sdl.QuitEvent:
            return true
        case *sdl.MouseMotionEvent:
            fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
                       t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
        }
    }
    return false
}