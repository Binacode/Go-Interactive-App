package main

import "fmt"

func main(){

		// for i := 0; i < 9; i++{
		// 	if i == 4{
		// 		break
		// 	}
		// 	fmt.Println(i)
		// }

		// for j := 5; j > 0; j--{
		// 	if j == 3{
		// 		continue
		// 	}
		// 	fmt.Println(j)
		// } 

		matriz := [3][3]int{}
		valor := 0
	
		for x := 0; x < 3; x++{
				for y := 0; y < 3; y++{
						matriz[x][y] = valor
						fmt.Print(matriz[x][y])
						valor++
				}
				fmt.Println()
		}

		fmt.Println(matriz)

}