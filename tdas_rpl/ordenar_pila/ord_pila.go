package ordenarpila

import TDAPila "tdas/pila"

// Implementar una función que ordene de manera ascendente una pila de enteros sin conocer su estructura interna
// y utilizando como estructura auxiliar sólo otra pila auxiliar.
// Por ejemplo, la pila [ 4, 1, 5, 2, 3 ] debe quedar como [ 1, 2, 3, 4, 5 ]
// (siendo el último elemento el tope de la pila, en ambos casos).
// Indicar y justificar el orden de la función.

func Ordenar(pila TDAPila.Pila[int]) {
	auxPila := TDAPila.CrearPilaDinamica[int]()

	for !pila.EstaVacia() {
		// Obtener el elemento actual de la pila
		elemento := pila.Desapilar()

		// Desapilar elementos de la pila auxiliar hasta encontrar el lugar adecuado para el elemento actual
		for !auxPila.EstaVacia() && auxPila.VerTope() < elemento {
			pila.Apilar(auxPila.Desapilar())
		}

		// Apilar el elemento actual en la posición correcta en la pila auxiliar
		auxPila.Apilar(elemento)
	}

	// Mover los elementos ordenados de la pila auxiliar a la pila original
	for !auxPila.EstaVacia() {
		pila.Apilar(auxPila.Desapilar())
	}
}
