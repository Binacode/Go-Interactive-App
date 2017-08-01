package main

import 	(
	"fmt"

	"github.com/vyr0n/Go-Interactive-App/06/paquetes/despedida"
	"github.com/vyr0n/Go-Interactive-App/06/paquetes/saludar"
)

func main(){
		nombre := "gente del futuro"

		saludar.Saludar(nombre)

		saludar.MeVes = "Esto es un texto asignado desde el main"
		fmt.Println(saludar.MeVes)

		despedida.Despedirse(nombre)

		// nombre := "gente del futuro"
		// saludar.Saludar(nombre)
		//

}
