package main

import (
    "fmt"
    "time"
)

func main()  {
    var h string
    go MostrarNumeros()
    fmt.Print("Digita algo ")
    fmt.Scan(&h)
    fmt.Println("Digitaste", h)
}

func MostrarNumeros()  {
    for i := 0; i <= 10; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}
