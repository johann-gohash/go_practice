package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    files, err := os.ReadDir("filedir")
    
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        fmt.Println(f.Name())

    }
}
