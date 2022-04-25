package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
    "time"
)

func main() {
    arguments := os.Args

    if len(arguments) == 1 {
        fmt.Println("Please provide port number")
        return
    }

    PORT :=  ":" + arguments[1]

    /*
        The net.listen() function makes the program a TCP server. 
        This function returns a Listener variable, which is a generic
        network listener for stream-oriented protocols.
    */
    l, err := net.Listen("tcp", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer l.Close()

    /*
        It is only after a successful call to Accept() that the TCP server
        begins to interact with TCP clients.

        The current implementation of the TCP server can only serve the first
        TCP client that connects to it, because the Accept() call is outside
        of the for loop. In the Create a Concurrent TCP Server section of this guide,
        you will see a TCP server implementation that can serve multiple TCP clients
        using Goroutines.
    */
    c, err := l.Accept()

    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        netData, err := bufio.NewReader(c).ReadString('\n')

        if err != nil {
            fmt.Println(err)
            return
        }

        if strings.TrimSpace(string(netData)) == "STOP" {
            fmt.Println("Exiting TCP server!")
            return
        }

        fmt.Print("-> ", string(netData))
        t := time.Now()
        
        myTime := t.Format(time.RFC3339) + "\n"
        c.Write([]byte(myTime))
    }



}
