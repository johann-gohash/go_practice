package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    // Get a greeting message and print in.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message
    message, err := greetings.Hello("Lise")

    // If an error was returned, print it to the console and
    // exit the program.
    if err  != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
