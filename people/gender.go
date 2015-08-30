package people

type Gender int

import (
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

const (
	MALE Gender = iota
	FEMALE
	NONBINARY
)

func (g Gender) String() string {
    if g == FEMALE {
        return "female"
    } else if g == MALE {
        return "male"
    } else {
        return "non-binary"
    }
}

func (g Gender) NameDirectory() string {
    if g == FEMALE {
        return "female"
    } else if g == MALE {
        return "male"
    } else if rand.Intn(99) + 1 < 50 {
        return "male"
    }
    return "female"
}
