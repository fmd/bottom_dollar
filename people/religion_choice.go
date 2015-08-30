package people

import (
    "fmt"
    "errors"
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
    c := rand.Intn(len(keyChoices) - 1)
    return ReligionChoiceFromKey(keyChoices[c])
}