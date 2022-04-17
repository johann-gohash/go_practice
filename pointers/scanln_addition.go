package main

import (
    "fmt"
)

func main () {
    var a int
    var b int

    // Scanln takes the address of a variable, and stores the user input there.
    // What happens if we don't provide the address? EDIT: It doesn't work.
    /*
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    */

    fmt.Scanln(a)
    fmt.Scanln(b)
    fmt.Println(a+b)
}
