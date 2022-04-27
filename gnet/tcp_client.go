package main

import (
    "log"
    "github.com/panjf2000/gnet"
)


func main() {
    
    log.Println("Hallooo")

    // NewClient needs a gnet.EventHandler
    v := gnet.EventHandler()
    x := gnet.NewClient()
}
