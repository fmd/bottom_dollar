package main

import (
    "encoding/json"
    "math/rand"
    "io/ioutil"
)

type BackgroundChoice struct {
    Name                  string   `json:"name"`
    ColorRange            []string `json:"color_range"`
    NaturalizationPercent int      `json:"naturalization_percent"`
    ImmigrationPercent    int      `json:"immigration_percent"`
    HasMiddleName         bool     `json:"has_middle_name"`
    Religions             []string `json:"religions"`
}

func AllBackgroundChoices() []BackgroundChoice {
    content, err := ioutil.ReadFile("backgrounds/backgrounds.json")
    if err != nil {
        panic(err)
    }

    var choices []BackgroundChoice
    err = json.Unmarshal(content, &choices)
    if err != nil{
        panic(err)
    }

    return choices
}

func BackgroundChoiceFromKey(key string) BackgroundChoice {
    choices := AllBackgroundChoices()
    for j := range choices {
        if key == choices[j].Name {
            return choices[j]
        }
    }

    panic("Could not find background.")
    return choices[0]
}

func BackgroundChoiceFromList(keyChoices []string) BackgroundChoice {
    c := rand.Intn(len(keyChoices) - 1)
    return BackgroundChoiceFromKey(keyChoices[c])
}
