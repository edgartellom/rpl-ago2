package ejercicio3

import (
	TDAPila "tdas/pila"
)

// Implementar una función que reciba un arreglo genérico e invierta su orden,
// utilizando los TDAs vistos. Indicar y justificar el orden de ejecución.

func InvertirArreglo[T any](arr []T) {
	pila := TDAPila.CrearPilaDinamica[T]()
	for _, elem := range arr {
		pila.Apilar(elem)
	}
	for i := 0; i < len(arr) && !pila.EstaVacia(); i++ {
		arr[i] = pila.Desapilar()
	}
}
