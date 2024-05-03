package main

import (
	variables "Usuario/go/Variables"
	"fmt"
	
)

func main() {

	//variables.MuestroEnteros()

	//variables.RestoVariable()

	estado,texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}	