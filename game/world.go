package game

import (
    "github.com/fmd/bronson"
)

type World struct {
    Place *Place
    Geometry *bronson.Geometry
}

func NewWorld() *World {
    w := &World{}
    opts := bronson.GeometryOpts{
        Width: 0.5,
        Height: 0.5,
    }

    g := bronson.NewGeometry(opts)
    w.Geometry = g
    return w
}