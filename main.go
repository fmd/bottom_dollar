package main

import (
	"github.com/fmd/bottom_dollar/graphics"
)

func main() {
    w := graphics.NewWindow()
    var shouldQuit bool
    for !shouldQuit {
        shouldQuit = ProcessOneFrameOfInput()
        w.ProcessOneFrame()
    }

    w.Destroy()
}