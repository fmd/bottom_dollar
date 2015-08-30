package main

import (
    "math/rand"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type Background struct {
    Name          string
    Color         string
    Immigrated    bool
    Naturalized   bool
    HasMiddleName bool
    Religion      Religion
}

func NewBackgroundFromChoice(choice BackgroundChoice) Background {
    b := &Background{}
    b.Name = choice.Name
    b.Color = choice.ColorRange[0] //TODO
    if (rand.Intn(99) )
    b.Immigrated
}