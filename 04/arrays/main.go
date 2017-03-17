package main

import "fmt"

func main(){
		// Arrays
		// Todos los valores deben ser del mismo tipo de dato.
		// Tamaño fijo

		// var names[4]string
		// names[0] = "Beli"
		// names[1] = "JJ"
		// names[2] = "Eli"
		// names[3] = "Fer"

		names := [4]string{"Beli", "JJ", "Eli", "Fer"}

		fmt.Println(names[0])
		size := len(names)
		fmt.Println("Size of array is: ", size)
		fmt.Println(names)
		// fmt.Println(names[inicio:final(excluyente)])
		fmt.Println(names[0:2])

}