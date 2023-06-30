package filtrar

import (
	"fmt"

	TDACola "tdas/cola"
)

func filtro(elem int) bool {
	return elem < 3
}

func EjecutarPruebas() {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 5; i++ {
		cola.Encolar(i)
	}
	FiltrarCola(cola, filtro)
	for !cola.EstaVacia() {
		fmt.Printf("%v ", cola.Desencolar())
	}
}
