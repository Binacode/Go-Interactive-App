package main

import "fmt"

func main(){
	// var n1, n2 int8
	// n1 = 15
	// n2 = 12
	// r := suma(n1, n2)

	// fmt.Println(r)

	/*
		var edad uint8
		edad = 3
		fmt.Println(tipoEdad(edad))
		edad = 15
		fmt.Println(tipoEdad(edad))
		edad = 21
		fmt.Println(tipoEdad(edad))
	*/

	n := []int8{52,63,47,5,5,3,7,6,100,2,47,-5}
	maximo, minimo := maxymin(n)
	fmt.Println("Máximo: ",maximo)
	fmt.Println("Minimo: ",minimo)
}

func suma(a, b int8)int8{
	return a + b
}

func tipoEdad(edad uint8)string{
		var tipo string
		switch{
			case edad < 12:
				tipo = "Niñez"
			case edad < 18:
				tipo = "Adolescencia"
			default:
				tipo = "Madurez"
		}
		return tipo
}


// func maxymin (numeros []int8) (int8, int8){
// Así era antes de declarar las variables.
func maxymin (numeros []int8) (max int8, min int8){
		// var max, min int8
		for _, v := range numeros{
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		// return max, min
		return
}