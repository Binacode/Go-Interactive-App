package main

import "fmt"

func main(){
	/*
		languages := make(map[string]string)
		languages["es"] = "Español"
		languages["en"] = "English"
		languages["fr"] = "Francias"

		fmt.Println(languages)
	*/

	/*
	idiomas := map[string]string{
			"es": "Español", 
			"en": "English", 
			"fr": "Francias",
			"pt": "Portuguese",
	}
	*/

	/* 
		 fmt.Println(idiomas)
  	 delete(idiomas, "en")
		 fmt.Println(idiomas)
  */
	
		// if idioma, ok := idiomas["pt"]; ok {
		// 	fmt.Println("Position exists an is ",idioma)
		// }else{
		// 	fmt.Println("Position doesn't exists")
		// }

	// fmt.Println(idiomas["pt"])

	numeros := map[int8]string{
		1: "This is uno",
		2: "This is dos",
		20: "This is other num",
		-4: "This is negative number",
	}


	fmt.Println(numeros)

}