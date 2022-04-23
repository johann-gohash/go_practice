package main

import (
    "fmt"
    "strings"
    //"reflect"
)

type User struct {
    Name string
    Pass string
}

func main() {

    the_string := "Dima:Pass123"
    y := strings.Split(*&the_string, ":")
    // fmt.Println(y)
    // fmt.Print(reflect.TypeOf(y))
    z := User{y[0], y[1]}
    fmt.Println(z)
}
