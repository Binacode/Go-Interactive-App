package main

import "fmt"

func main(){

		//Slices
		//1. Apuntador a un array
		//2. Tamaño (no es fijo)
		//3. Capacidad

		// var names []string
		// make (tipo de dato del slice, tamaño inicial, capacidad inicial)
		
		// names := make([]string,0)
		/*
			fmt.Printf("Size: %d and capacity: %d\n", len(names), cap(names))
			names = append(names, "JJ")
			fmt.Println(names)
			fmt.Printf("Size: %d and capacity: %d\n", len(names), cap(names))
			names = append(names, "Beli")
			fmt.Println(names)
			fmt.Printf("Size: %d and capacity: %d\n", len(names), cap(names))
			names = append(names, "Eli")
			fmt.Println(names)
			fmt.Printf("Size: %d and capacity: %d\n", len(names), cap(names))
			names = append(names, "Naoli")
			fmt.Println(names)
			fmt.Printf("Size: %d and capacity: %d\n", len(names), cap(names))
			names = append(names, "Melqui")
			fmt.Println(names)
			fmt.Printf("Size: %d and capacity: %d\n", len(names), cap(names))
			names = append(names, "Fer")
			fmt.Println(names)
			fmt.Printf("Size: %d and capacity: %d\n", len(names), cap(names))
		*/

		names := []string{
			"elle",
			"bell",
			"naol",
			"fern",
		}

		fmt.Println(names)

}