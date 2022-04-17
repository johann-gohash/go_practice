package main

import "fmt"

func main() {

    x := 7
    y := &x

    *y += 2
    fmt.Println(x)

}
