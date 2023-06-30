package claves

import TDALista "tdas/lista"

// Implementar una primitiva para el ABB, que devuelva una lista con las claves del mismo,
// ordenadas tal que si insertáramos las claves en un ABB vacío, dicho ABB tendría la misma estructura que el árbol original.

func (abb *abb[K, V]) Claves() TDALista.Lista[K] {
	lista := TDALista.CrearListaEnlazada[K]()
	abb.raiz.claves(lista)
	return lista
}

func (nodo *nodoAbb[K, V]) claves(lista TDALista.Lista[K]) {
	if nodo == nil {
		return
	}
	lista.InsertarUltimo(nodo.clave)
	nodo.izquierdo.claves(lista)
	nodo.derecho.claves(lista)
}
