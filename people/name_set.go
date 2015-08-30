package people

import (
    "fmt"
    "io/ioutil"
    "math/rand"
)

type NameSet []string

func NameSetFromList(list []NameSet) NameSet {
    index := rand.Intn(len(list)) - 1
    return list[index]
}

func NameList(t string, g string, b string) []NameSet {
    var content []byte
    var err error

    if t == "last" {
        content, err = ioutil.ReadFile(fmt.Sprintf("names/%s/%s.json", t, b))
    } else {
        content, err = ioutil.ReadFile(fmt.Sprintf("names/%s/%s/%s.json", t, g, b))
    }

    if err != nil {
        panic(err)
    }

    var choices []NameSet
    err = json.Unmarshal(content, &choices)
    if err != nil {
        panic(err)
    }

    return choices
}
