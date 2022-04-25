package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main () {

    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide host:port")
        return
    }

    CONNECT := arguments[1]

    /*
        A call to net.Dial() begins the implementation of the TCP client
        and will connect you to the desired  TCP server.

        The second parameter of net.Dial()has two parts; the first is the hostname
        or the IP address of the TCP server and the second is the port number the TCP
        server listens on.
    */

    c, err := net.Dial("tcp", CONNECT)

    if err != nil {
        fmt.Println(err)
        return
    }


    /*
        The entire for loop that is used to read user input will only terminate when you
        send the STOP commandd to the TCP server.
    */
    for {
        /*
            bufio reader and the bufio.NewReader(c).ReadString('\n') statement read
            the TCP server's response. The error vrariaable is ignored here for simplicity.
        */
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> ")

        text, _ := reader.ReadString('\n')
        fmt.Fprintf(c, text+"\n")

        message, _ := bufio.NewReader(c).ReadString('\n')
        fmt.Print("->: " + message)

        if strings.TrimSpace(string(text)) == "STOP" {
            fmt.Println("TCP client exiting...")
            return
        }
    }

}
