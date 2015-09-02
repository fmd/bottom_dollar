package main

import (
    "github.com/fmd/bottom_dollar/graphics"
    "github.com/fmd/bottom_dollar/places"
    "fmt"
)

func main() {
    w := NewWindow()
    var shouldQuit bool

    world := NewWorld()
    world.Place = MakeBar()

    for !shouldQuit {
        shouldQuit = ProcessOneFrameOfInput()
        ProcessOneFrameOfRendering()
    }

    w.Destroy()
}

func ProcessOneFrameOfRendering() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    graphics.RenderAllRenderables()

    sdl.GL_SwapWindow(w.Window)
}

func MakeBar() *Place {
    p := &places.Place{}
    p.Layers = append(p.Layers, places.NewExampleGroundLayer())
    p.Layers = append(p.Layers, places.NewExampleWallLayer())
    return p
}