package people

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "time"
    "math/rand"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type NameSet []string

func NameSetFromList(list []NameSet) NameSet {
    index := rand.Intn(len(list))
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
