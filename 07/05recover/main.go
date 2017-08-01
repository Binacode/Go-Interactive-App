package main

import "fmt"

func main()  {
    f()
}

func f()  {
    defer func ()  {
        if r := recover(); r != nil {
            fmt.Println("%T\n", r)
            fmt.Println("Recuperado en f: ",r)
        }
    }()
    fmt.Println("Llamando a g.")
    g(4)
    fmt.Println("Retorna normalmente desde g")
}

func g(i int)  {
    if i > 3{
        fmt.Println("AAAAAAH (entro en panico)")
        panic("El número no puede ser mayor que 3")
    }
    defer fmt.Println("Defer en la funcion g")
    fmt.Println("Imprimiendo en g", i)
}
