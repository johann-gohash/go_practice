package main

import (
    "log"
    "fmt"
)

func main () {
    v := "stringggg"
    x := 3
    y := 3.141592
    z := 2.718


    log.Println("Literal", v, "literal", x, "literal", y, "zzz", z)

    // Stick with this for concatenating strings.
    // the S---- variants return strings instead of just printing them to stdout.
    a := fmt.Sprint("Literal", v, "love is a verb.")
    fmt.Print(a)
}
