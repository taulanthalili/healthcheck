package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)


var NGINX_LOG_PATH = "../access.log"

func main() {
    var text string

    fmt.Print("Enter text: ")
    // get the sub string to search from the user
    fmt.Scanln(&text)

    // read the whole file at once
    b, err := ioutil.ReadFile(NGINX_LOG_PATH)
    if err != nil {
        panic(err)
    }
    s := string(b)
    // //check whether s contains substring text
    fmt.Println(strings.Contains(s, text))
}