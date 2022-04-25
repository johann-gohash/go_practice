package main

import (
    "fmt"
    "encoding/json"
)

type Book struct {
    Title string `json:"title"`
    Author Author `json:"author"`
}

type Author struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Developer bool `json:"is_developer"`
}

func marshalling ()  {

    sample_author := Author{Name: "Elliot Forbes", Age: 25, Developer: true}
    book := Book{"Learning Concurrency in Python", sample_author}

    // fmt.Printf("%+v\n", book)

    byteArray, err := json.MarshalIndent(book, "", "  ")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(byteArray))
}

func main () {
    marshalling()
}
