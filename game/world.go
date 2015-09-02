package game

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/fmd/bottom_dollar/engine"
)

type World struct {
    Place *Place
}

func NewWorld() *World {
    w := &World{}
    engine.AddToRenderList(w)
    return w
}

func (w *World) Render() {
    gl.Begin(gl.TRIANGLES)
    gl.Color3f(1.0, 0.0, 0.0)
    gl.Vertex2f(0.5, 0.0)
    gl.Color3f(0.0, 1.0, 0.0)
    gl.Vertex2f(-0.5, -0.5)
    gl.Color3f(0.0, 0.0, 1.0)
    gl.Vertex2f(-0.5, 0.5)
    gl.End()
}