package people

import (
	"math/rand"
	"time"
)

type Gender int

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
	} else if rand.Intn(100) < 50 {
		return "male"
	}
	return "female"
}

func RandomGender() Gender {
	choice := rand.Intn(100)
	if choice < 50 {
		return FEMALE
	}

	if choice < 99 {
		return MALE
	}

	return NONBINARY
}
