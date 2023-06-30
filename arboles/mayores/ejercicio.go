package mayores

import TDALista "tdas/lista"

// Se tiene un árbol binario de búsqueda con cadenas como claves y función de comparación strcmp.
// Implementar una primitiva func (abb *abb[K, V]) Mayores(clave K) Lista[K] que, dados un ABB y una clave,
// devuelva una lista ordenada con las claves del árbol estrictamente mayores a la recibida por parámetro (que no necesariamente está en el árbol).
// Implementar sin utilizar el iterador Interno del ABB.

func (abb *abb[K, V]) Mayores(clave K) TDALista.Lista[K] {
	mayores := TDALista.CrearListaEnlazada[K]()

	abb.raiz.abbMayores(clave, abb.cmp, mayores)
	return mayores
}

func (nodo *nodoAbb[K, V]) abbMayores(clave K, cmp funcCmp[K], claves TDALista.Lista[K]) {
	if nodo == nil {
		return
	}
	if cmp(nodo.clave, clave) > 0 {
		nodo.izquierdo.abbMayores(clave, cmp, claves)
		claves.InsertarUltimo(nodo.clave)
	}
	nodo.derecho.abbMayores(clave, cmp, claves)
}
