package ejercicios

import (
	"fmt"
	"strconv"
)

func Ejercicio01(numero string) (int, string) {
	numeroEntero, err := strconv.Atoi(numero)

	if err != nil {
		return 0, "El valor ingresado no es un número válido"
	}

	if numeroEntero > 100 {
		fmt.Println(numero + "\n", "El numero es mayor a 100")
	} else {
		fmt.Println( numero + "\n","El numero es menor a 100")
	}

	return numeroEntero, ""
}
