package main


import (
    "errors"
    "fmt"
    "os"
)


func main() {
    _, err := os.Stat("file.txt")

    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File doesn't exist")
    } else {
        fmt.Println("File exists.")
    }

}
