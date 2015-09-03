package game

import (
)

type World struct {
    Place *Place
}

func NewWorld() *World {
    w := &World{}
    return w
}