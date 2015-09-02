package main

import (
    "github.com/fmd/bottom_dollar/engine"
    "github.com/fmd/bottom_dollar/game"
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
)

var shouldQuit bool

func main() {
    e := engine.NewEngine(engine.WindowOpts{Title: "Bottom Dollar", Width: 800, Height: 600})
    w := game.NewWorld()
    fmt.Println(w)
    mainLoop(e)
}

func mainLoop(e *engine.Engine) {
    shouldQuit = false
    for !shouldQuit {
        handleInput(e.ReceiveEvents())
        e.ProcessOneFrame()
    }
}

func handleInput(b engine.EventBuffer) {
    for _, event := range b {
        switch t := event.(type) {
            case *sdl.QuitEvent:
                shouldQuit = true
            case *sdl.MouseMotionEvent:
                fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
                           t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
        }
    }
}

func MakeBar() *game.Place {
    p := &game.Place{}
    p.Layers = append(p.Layers, game.NewExampleGroundLayer())
    p.Layers = append(p.Layers, game.NewExampleWallLayer())
    return p
}