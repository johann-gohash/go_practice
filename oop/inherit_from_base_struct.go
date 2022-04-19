package main

import (
    "fmt"
)

type BaseStruct struct {
    integer_data int  // private
}

type SuperStruct struct {
    BaseStruct
    integer_data2 int

}

func (bs BaseStruct) SayHello() {
    fmt.Println("Hello")
}

func main() {

    x := BaseStruct{2}
    x.SayHello()

    y := SuperStruct{BaseStruct{3}, 4}
    y.SayHello()
}

