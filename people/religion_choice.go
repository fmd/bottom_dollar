package people

import (
	"encoding/json"
	"errors"
	"fmt"
    "time"
	"io/ioutil"
	"math/rand"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type ReligionChoice struct {
	Name             string `json:"name"`
	Plural           string `json:"plural"`
	ReligiousPercent int    `json:"religious_percent"`
}

func AllReligionChoices() []ReligionChoice {
	content, err := ioutil.ReadFile("backgrounds/religions.json")
	if err != nil {
		panic(err)
	}

	var choices []ReligionChoice
	err = json.Unmarshal(content, &choices)
	if err != nil {
		panic(err)
	}

	return choices
}

func ReligionChoiceFromKey(key string) (ReligionChoice, error) {
	choices := AllReligionChoices()
	for j := range choices {
		if key == choices[j].Name {
			return choices[j], nil
		}
	}

	return ReligionChoice{}, errors.New(fmt.Sprintf("Could not find religion: %s", key))
}

func ReligionChoiceFromList(keyChoices []string) (ReligionChoice, error) {
	c := rand.Intn(len(keyChoices))
	return ReligionChoiceFromKey(keyChoices[c])
}
