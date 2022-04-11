package main

import (
    "fmt"
    "os"
    "strings"
    "reflect"
)

func main () {
    fmt.Println(strings.Join(os.Args[1:], "  "))
    /*
        Question to self:
            Is os.Args[1:] an iterator? How it's used in a range for loop
            seems to suggest so.

        Answer: 
            Nope. os.Args[1:] returns a list, and range ITERATES over a list.

    */
    // fmt.Println(strings.Join(os.Args[1:], "  "))
    fmt.Println(os.Args[1:])
    fmt.Println(reflect.TypeOf(os.Args[1:]))
}
