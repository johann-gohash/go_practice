package main

import (
    "fmt"
    "log"
    "os"
)

func main() {

    fInfo, err := os.Stat("file.txt")

    if err != nil {
        log.Fatal(err)
    }

    fsize := fInfo.Size()

    fmt.Printf("The file size is %d bytes\n", fsize)
}

