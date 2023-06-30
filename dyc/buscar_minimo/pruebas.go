package minimo

import "fmt"

func EjecutarPruebas() {
	arr := []int{8, 3, 6, 2, 7, 1, 9, 5, 4}
	fmt.Println("Minimo:", BuscarMinimo(arr))
}
