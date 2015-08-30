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

type BackgroundChoice struct {
	Name                  string   `json:"name"`
	ColorRange            []string `json:"color_range"`
	NaturalizationPercent int      `json:"naturalization_percent"`
	ImmigrationPercent    int      `json:"immigration_percent"`
	HasMiddleName         bool     `json:"has_middle_name"`
	AllowLastNameChange   bool     `json:"allow_last_name_change"`
	NameChoices           []string `json:"name_choices"`
	Religions             []string `json:"religions"`
}

func AllBackgroundChoices() []BackgroundChoice {
	content, err := ioutil.ReadFile("backgrounds/backgrounds.json")
	if err != nil {
		panic(err)
	}

	var choices []BackgroundChoice
	err = json.Unmarshal(content, &choices)
	if err != nil {
		panic(err)
	}

	return choices
}

func BackgroundChoiceFromKey(key string) (BackgroundChoice, error) {
	choices := AllBackgroundChoices()
	for j := range choices {
		if key == choices[j].Name {
			return choices[j], nil
		}
	}
	return BackgroundChoice{}, errors.New(fmt.Sprintf("Could not find background: %s", key))
}

func BackgroundChoiceFromList(keyChoices []string) (BackgroundChoice, error) {
	c := rand.Intn(len(keyChoices))
	return BackgroundChoiceFromKey(keyChoices[c])
}

func RandomBackgroundChoice() BackgroundChoice {
	backgrounds := AllBackgroundChoices()
	c := rand.Intn(len(backgrounds))
	return backgrounds[c]
}
