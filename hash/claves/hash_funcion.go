package claves

import TDALista "tdas/lista"

// Implementar una función func Claves[K comparable, V any](Diccionario[K, V]) Lista[K]
// que reciba un diccionario y devuelva una lista con sus claves.

func Claves[K comparable, V any](dic Diccionario[K, V]) TDALista.Lista[K] {
	lista := TDALista.CrearListaEnlazada[K]()
	for iter := dic.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		lista.InsertarUltimo(clave)
	}
	return lista
}
