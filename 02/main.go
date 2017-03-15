package main

import "fmt"

func main() {
	// const nombre = "Jonathan"
	// fmt.Println(nombre)
	var a int 
	var b int8 
	n := "Pedro"
	p := "Juarez"
	a = 121212
	b = 5

	// casting
	
	c:= a + int(b)
	fmt.Println(c)
	// fmt.Printf("Hola," + n + "Cómo andas?")
	fmt.Printf("Hola, %s %s %d Cómo andas?\n", n, p, b)
	fmt.Printf("c es de tipo: %T\n", c)

	// Prioridad aritmética
	// * / + -
	// () % * / + -
	// módulo es el residuo de la división.
	// 14 % 6
	// d := 6 + 6*6 - 6
	// fmt.Println(d)
	var in int8
	var nombre string
	var numero int
	var aprender bool
	fmt.Printf(nombre, numero, aprender)
	
	in = 14 % 6
	fmt.Printf("\n numero: %d",in)
}
