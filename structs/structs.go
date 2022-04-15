package main

import (
    "fmt"
)

func main() {

    fmt.Printf("Working")

    type ninja struct {
        name string
        weapons []string
        level int
    }

    wallace := ninja{name: "Wallace"}

    wallace = ninja {"Wallace", []string{"M-16", "M1911", "Knife", "Flashbang"},1}

    fmt.Println(wallace)
    fmt.Println(wallace.name)
    fmt.Println(wallace.weapons)
    fmt.Println(wallace.level)

    wallace.level++
    wallace.weapons = append(wallace.weapons, "RPG-7")

    type dojo struct {
        name string 
        ninja ninja
    }

    golangDojo := dojo {
        name: "Golang Dojo",
        ninja: wallace,
    }

    fmt.Println(golangDojo)
    fmt.Println(golangDojo.ninja.level)
    golangDojo.ninja.level = 4

    fmt.Println(golangDojo.ninja.level)
    fmt.Println(wallace.level)


    type addressedDojo struct {
        name string
        ninja *ninja
    }

    addressed := addressedDojo{"Addressed Go Dojo", &wallace}
    fmt.Println(*addressed.ninja)
    addressed.ninja.level = 5

    fmt.Println(addressed.ninja.level)
    fmt.Println(wallace.level)
    
// "New" way to instantiate a struct in Go
    johnnyPointer := new(ninja)
    fmt.Println(johnnyPointer)
    fmt.Println(*johnnyPointer)
    johnnyPointer.name = "Johnny"
    johnnyPointer.weapons = []string{"Ninja Star"}
    johnnyPointer.level = 1
    fmt.Println(*johnnyPointer)

    intern := ninjaIntern{name: "Lise", level: 1}
    intern.sayHello()
    intern.sayName()
}


type ninjaIntern struct {
    name string
    level int
}

func (ninjaIntern) sayHello() {
    fmt.Println("hello")
}

func (n ninjaIntern) sayName() {
    fmt.Println(n.name)
}
