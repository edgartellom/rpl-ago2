package claves

import TDALista "tdas/lista"

// Para un hash abierto, implementar una primitiva func (hash *hashCerrado[K, V]) Claves() Lista[K]
// que devuelva una lista con sus claves, sin utilizar el iterador interno.

func (hash *hashAbierto[K, V]) Claves() TDALista.Lista[K] {
	listaRes := TDALista.CrearListaEnlazada[K]()
	for _, lista := range hash.tabla {
		for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			par := iter.VerActual()
			listaRes.InsertarUltimo(par.clave)
		}
	}
	return listaRes
}
