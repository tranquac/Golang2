package main

import (
    "fmt"
    //"runtime"
    "sync"
)

var wg sync.WaitGroup

func main() {
    // add one thing to wait for
    wg.Add(1)
    go woof()
    bar()

    // stop waiting when 0 things in waiting list
    wg.Wait()
}

func woof() {
    for i := 1; i <= 5; i++ {
        fmt.Println("woof ", i)
    }
    // remove one thing from waiting list
    wg.Done()
}

func bar() {
    for i := 1; i <= 5; i++ {
        fmt.Println("meow ", i)
    }
}