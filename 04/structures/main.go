package main

import "fmt"

//Persona es una estructura
type Persona struct{
	Nombre string
	Edad 	 uint8
	Emails []string
}

func main(){

/*
	var persona1 Persona
	persona1.Nombre = "JJ"
	persona1.Edad = 25
	fmt.Println(persona1)
	fmt.Println(persona1.Nombre)
	fmt.Println(persona1.Edad)
*/
		emails := []string{"mail-1@b.com","mail-2@b.com"}

		persona2 := Persona{
			// Nombre: "Andy",
			// Edad: 23,
			"Andy",
			23,
			// []string{"mail-1@b.com","mail-2@b.com"},
			emails,
		}
		fmt.Println(persona2)



}