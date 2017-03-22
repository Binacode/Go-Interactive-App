package main

import "fmt"

func main(){
	/*
		numeros := []string{"cero", "uno", "dos", "tres"}

		for i, v := range numeros{
			fmt.Println(i, v)
		}

		for _, v := range numeros{
			fmt.Println(v)
		}

		for i := range numeros{
			fmt.Println(i)
		}
	*/
	
		// nombres := map[string]string {	"es": "España", "co": "Colombia", "br": "Brasil"	}

		// for i,v := range nombres {
		// 	fmt.Println(i, v)
		// }
		// for _,v := range nombres {
		// 	fmt.Println(v)
		// }
		// for i := range nombres {
		// 	fmt.Println(i)
		// }

		frase := "Hola mundo"
		
		for posicion, letra := range frase{
				fmt.Println(posicion, string(letra))
		}

		// for i := 32; i < 200; i++{
		// 	fmt.Println(string(i))
		// 

		for _, entero := range []int{15, 36, 24, 85}{
			fmt.Println(entero)
		}
		for _, entero := range "This is an example"{
			fmt.Println(string(entero))
		}
}