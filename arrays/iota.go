package main


import (
    "fmt"
)

func main () {

    type Currency int

    const (
        USD Currency = iota 
        EUR
        GBP
        RMB
    )
}
