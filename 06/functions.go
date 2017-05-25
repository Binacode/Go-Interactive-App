package main

import "fmt"


func main(){

	saludar("Alexia", 31)

}

func saludar(nombre string, edad uint8){
		// fmt.Println("Hi!, ", nombre, "tienes ", edad)
		fmt.Printf("Hola %s tienes %d años de edad\n", nombre, edad)
		if edad > 30{
			return
		}
		fmt.Println("You're young.")
}
