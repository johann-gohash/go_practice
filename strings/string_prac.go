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


    y := strings.Split("Dima:Pass123", ":")
    // fmt.Println(y)
    // fmt.Print(reflect.TypeOf(y))
    z := User{y[0], y[1]}
    fmt.Println(z)


}
