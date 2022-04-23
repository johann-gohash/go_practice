package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    content, err := ioutil.ReadFile("filedir/file_0.txt")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(content))
    fmt.Printf("%T\n", content)
}
