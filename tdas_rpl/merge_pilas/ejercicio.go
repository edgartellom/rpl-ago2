package mergepilas

import TDAPila "tdas/pila"

// Dadas dos pilas de enteros positivos (con posibles valores repetidos) cuyos elementos fueron ingresados de menor a mayor,
// se pide implementar una función func MergePilas(pila1, pila2 Pila[int]) []int que devuelva un array ordenado de menor a mayor
// con todos los valores de ambas pilas sin repeticiones.
// Detallar y justificar la complejidad del algoritmo considerando que el tamaño de las pilas es N y M respectivamente.

func MergePilas(pila1, pila2 TDAPila.Pila[int]) []int {
	aux := TDAPila.CrearPilaDinamica[int]()
	for !pila1.EstaVacia() && !pila2.EstaVacia() {
		var actual int
		if pila1.VerTope() >= pila2.VerTope() {
			actual = pila1.Desapilar()
		} else {
			actual = pila2.Desapilar()
		}
		apilarSinRepetir(aux, actual)
	}
	for !pila1.EstaVacia() {
		actual := pila1.Desapilar()
		apilarSinRepetir(aux, actual)
	}
	for !pila2.EstaVacia() {
		actual := pila2.Desapilar()
		apilarSinRepetir(aux, actual)
	}
	var res []int
	for !aux.EstaVacia() {
		res = append(res, aux.Desapilar())
	}
	return res
}

func apilarSinRepetir(pila TDAPila.Pila[int], dato int) {
	if pila.EstaVacia() || pila.VerTope() != dato {
		pila.Apilar(dato)
	}
}
