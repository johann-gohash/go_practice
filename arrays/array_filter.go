package main

import (
    "fmt"
    "strings"
)

func filter(ss []string, test func(string) bool) (ret [] string) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s) 
        }
    }
    return ret
}

func main () {

    ss := []string{"foo_1", "asdf", "looong", "nfoo_1", "foo_2"}

    mytest := func(s string) bool { return !strings.HasPrefix(s, "foo_") && len(s) <= 7 }
    s2 := filter(ss, mytest)

    fmt.Println(s2)

}
