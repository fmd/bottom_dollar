package people

import (
    "time"
    "math/rand"
)

type SocioEconomicClass int

func init() {
    rand.Seed(time.Now().UnixNano())
}

const (
    POOR SocioEconomicClass = iota
    WORKING
    MIDDLE
    RICH
)

func (s SocioEconomicClass) String() string {
    switch s {
    case POOR:
        return "poor"
    case WORKING:
        return "working class"
    case MIDDLE:
        return "middle class"
    case RICH:
        return "upper class"
    }
    return "nil"
}

func RandomSocioEconomicClassForBackground(b *Background) SocioEconomicClass {
    return WORKING
}