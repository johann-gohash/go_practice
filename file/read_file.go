package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    content, err := ioutil.ReadFile("file.txt")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(content))
    fmt.Printf("%T\n", content)
}
