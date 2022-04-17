package main

import (
    "fmt"
    "strings"
)


func capitalize(sample *string) {

    //fmt.Println((*sample).ToUpper())
    fmt.Println(strings.ToUpper(*sample))
    *sample = strings.ToUpper(*sample)
}

func main() {

    x := "woo"
    capitalize(&x)
    fmt.Println(x)
}
