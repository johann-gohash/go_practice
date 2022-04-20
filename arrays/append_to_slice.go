package main


import (
    "fmt"
    "time"
)

func main () {

    i := 0
    var sliceint []int
    for {
        sliceint = append(sliceint, i)
        time.Sleep(time.Second)
        fmt.Println(len(sliceint))
    }
}
