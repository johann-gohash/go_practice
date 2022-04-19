package main

import (
    "errors"
    "log"
)

var (
    ErrYoureFuckedUp = errors.New("You talk like a fag and your shit's all retarded.")
)


func EvenOrError (number int) error { 
    var err error
    if number%2 != 0 {
        err = ErrYoureFuckedUp
    }
    return err
}

func main () {
    x := EvenOrError(3)
    if x != nil {
        log.Fatal("Ain't even")
    }
}
