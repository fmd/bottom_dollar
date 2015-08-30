package main

import (
    "encoding/json"
    "math/rand"
    "io/ioutil"
)

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
    if err != nil{
        panic(err)
    }

    return choices
}

func ReligionChoiceFromKey(key string) ReligionChoice {
    choices := AllReligionChoices()
    for j := range choices {
        if key == choices[j].Name {
            return choices[j]
        }
    }

    panic("Could not find religion.")
    return choices[0]
}

func ReligionChoiceFromList(keyChoices []string) ReligionChoice {
    c := rand.Intn(len(keyChoices) - 1)
    return ReligionChoiceFromKey(keyChoices[c])
}