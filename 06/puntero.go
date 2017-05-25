package main

import "fmt"


func main(){

		a := 3
		// fmt.Println(&a)
		fmt.Println("Antes de duplicar, a = ",a)
		// duplicar(a)
		duplicar(&a)
		fmt.Println("Después de duplicar, a = ",a)
}


func duplicar(a *int){
	// *x = *x * 2
	*a *= 2
	fmt.Println("Dentro de la función deplicar vale", *a)
}