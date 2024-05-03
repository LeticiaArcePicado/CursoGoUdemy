package variables

import (
	"fmt"
	"time"
	"strconv"
)

var  Nombre string 
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariable (){
	
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.29
	Fecha = time.Now()

	fmt.Println("Nombre ", Nombre)
	fmt.Println("Estado ", Estado)
	fmt.Println("Sueldo ", Sueldo)
	fmt.Println("Fecha ", Fecha)
}

func ConviertoATexto(numero int) (bool,string) {
	var texto = strconv.Itoa(numero)
	return true ,texto
}