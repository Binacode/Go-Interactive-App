package main

import "fmt"

func main(){
		
		// var a int
		// a = 1

		// switch a{
		// 	case 1: 
		// 		fmt.Println("The value is 1")
		// 	case 2:
		// 		fmt.Println("The value is 2")
		// 	default:
		// 		fmt.Println("Isn't 1 and 2")
		// }


		/*
		 switch a{
			case 1: 
				fmt.Println("Monday")
			case 2:
				fmt.Println("Tuesday")
			case 3:
				fmt.Println("Wednesday")
			case 4:
				fmt.Println("Thursday")
			case 5:
				fmt.Println("Friday")
			case 6:
				fmt.Println("Saturday")
			case 7:
				fmt.Println("Sunday")
			default:
				fmt.Println("There isn't a day of the week.")
		}
		*/

		/*
		switch a {
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 5:
				fmt.Println("You are on week")
			case 6:
				fallthrough
			case 7:
				fmt.Println("You are on weekend :)")
			default:
				fmt.Println("It isn't a valid day")
		}*/
		
			switch a := 30; {
			case a>= 0 && a <= 5:
				fmt.Println("You are on week")
			case a>= 6 && a <= 7:
				fmt.Println("You are on weekend")
			default:
				fmt.Println("It isn't a valid day")
		}


}