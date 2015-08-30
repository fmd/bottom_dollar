package main

import (
    "math/rand"
)

type SocioEconomicClass int

const (
    HOMELESS SocioEconomicClass = iota
    WORKING
    LOWER_MIDDLE
    MIDDLE
    UPPER_MIDDLE
    UPPER
)

func (s SocioEconomicClass) String() string {
    switch s {
    case HOMELESS:
        return "homeless"
    case WORKING:
        return "working class"
    case MIDDLE:
        return "middle class"
    case UPPER:
        return "upper class"
    }
    return "middle class"
}

func RandomSocioEconomicClass() SocioEconomicClass {
    choice := rand.Intn(99)
    switch {
    case choice < 5:
        return HOMELESS
    case choice < 60:
        return WORKING
    case choice < 95:
        return MIDDLE
    case choice < 100:
        return UPPER
    }
    return WORKING
}