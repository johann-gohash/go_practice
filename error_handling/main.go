package main

import (
    "errors"
    "log"
)

var (
    ErrSomethingMessedUp = errors.New("There's an issue I refuse to elaborate.")
)


func EvenOrError (number int) error { 
    var err error
    if number%2 != 0 {
        err = ErrSomethingMessedUp
    }
    return err
}

func main () {
    x := EvenOrError(3)
    if x != nil {
        log.Fatal("Ain't even")
    }
}
