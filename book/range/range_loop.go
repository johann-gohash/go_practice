package main

import (
    "fmt"
)

func main() {

    /*
        In Go, a range loop is a variant of for-loop. Because
        every loop in Go is a for-loop.
        Small test: Use a range loop to iterate over an array of strings.
    */
   
    name_list := [4]string{"Ricardo", "Grey Ghost", "Hattie", "Rosemly"}

    for _, arg := range name_list {
        fmt.Println(arg)
    }
}
