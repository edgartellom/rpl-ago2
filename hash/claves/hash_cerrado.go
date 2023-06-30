package claves

import TDALista "tdas/lista"

// Para un hash cerrado, implementar una primitiva func (hash *hashCerrado[K, V]) Claves() Lista[K]
// que devuelva una lista con sus claves, sin utilizar el iterador interno.

func (hash *hashCerrado[K, V]) Claves() TDALista.Lista[K] {
	lista := TDALista.CrearListaEnlazada[K]()
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			lista.InsertarUltimo(celda.clave)
		}
	}
	return lista
}
