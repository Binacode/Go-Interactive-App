package main

import "fmt"

func main()  {
    division(3,0)
}

func division(dividendo, divisor int8)  {
    defer fmt.Println("Esto se ejecutará pase lo que pase")
    if divisor == 0 {
        panic("No se puede dividir por cero")
    }
    fmt.Println(dividendo / divisor)
}
