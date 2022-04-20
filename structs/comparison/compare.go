package main

import (
    "log"
)


type NamePass struct {
    name string
    pass string
}



func main () {
    a := NamePass{"bob", "pass123"}
    b := NamePass{"bob", "pass123"}

    if a == b {
        log.Println("Same lol")
    } else {
        log.Println("Not same ;-;")
    }
}

