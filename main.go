package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    res, err := http.Get("https://test15384.docs.apiary.io/#reference/0/group-user/register")
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    retJSON := struct {
        MyKey       string `json:"key"`
        Permissions []string
    }{"w", []string{"w", "w"}}

    w, err := json.Marshal(retJSON)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(w), retJSON)

    fmt.Println(string(body))

    if err := json.Unmarshal(body, &retJSON); err != nil {
        panic(err)
    }

    fmt.Println(retJSON)

}
