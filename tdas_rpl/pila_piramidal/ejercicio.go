package piramidal

import TDAPila "tdas/pila"

// Dada una pila de punteros a enteros, escribir una función que determine si es piramidal.
// Una pila de enteros es piramidal si cada elemento es menor a su elemento inferior
// (en el sentido que va desde el tope de la pila hacia el otro extremo). La pila no debe ser modificada.

func EsPiramidal(pila TDAPila.Pila[int]) bool {
	var backup []int
	var resul bool = true
	for !pila.EstaVacia() {
		tope := pila.Desapilar()
		backup = append(backup, tope)
		if !pila.EstaVacia() {
			inf := pila.VerTope()
			if !(tope < inf) {
				resul = false
			}
		}
	}
	for i := len(backup) - 1; i >= 0; i-- {
		pila.Apilar(backup[i])
	}
	return resul
}
