package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 10; i++ {
        fmt.Println(from, ":", i)
        time.Sleep(time.Second)
    }
}

func main(){
    f("direct")

    go f("goroutine")
    go f("alejandro")
    go f("caballero")
    go f("salas")

    go func(msg string){
        fmt.Println(msg)
    }("going")

    time.Sleep(77 * time.Second)
    fmt.Println("done")
}
