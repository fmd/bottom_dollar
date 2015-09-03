package main

import (
    "github.com/fmd/bronson"
    "github.com/fmd/bottom_dollar/game"
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
)

var shouldQuit bool

func main() {
    e := bronson.New(bronson.WindowOpts{Title: "Bottom Dollar", Width: 800, Height: 600})
    w := game.NewWorld()
    fmt.Println(w)
    mainLoop(e)
}

func mainLoop(e *bronson.Bronson) {
    shouldQuit = false
    for !shouldQuit {
        handleInput(e)
        e.ProcessOneFrame()
    }
}

func handleInput(b *bronson.Bronson) {
    for _, event := range b.ReceiveEvents() {
        switch /*t := */ event.(type) {
            case *sdl.QuitEvent:
                shouldQuit = true
                b.Cleanup()
                break
            case *sdl.MouseMotionEvent:
                break
                //fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
                //           t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
        }
    }
}

func MakeBar() *game.Place {
    p := &game.Place{}
    p.Layers = append(p.Layers, game.NewExampleGroundLayer())
    p.Layers = append(p.Layers, game.NewExampleWallLayer())
    return p
}