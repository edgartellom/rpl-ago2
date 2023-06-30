package ejercicio5

// Dada una lista enlazada implementada con las siguientes estructuras:

// type nodoLista[T any] struct {
//     prox *nodoLista[T]
//     dato T
// }

// type ListaEnlazada[T any] struct {
//     prim *nodoLista[T]
// }

// Escribir una primitiva que reciba una lista y devuelva el elemento que esté a k posiciones del final
// (el ante-k-último), recorriendo la lista una sola vez y sin usar estructuras auxiliares.
// Considerar que k es siempre menor al largo de la lista.
// Por ejemplo, si se recibe la lista [ 1, 5, 10, 3, 6, 8 ], y k = 4, debe devolver 10.
// Indicar el orden de complejidad de la primitiva.

func (lista *ListaEnlazada[T]) AnteKUltimo(k int) T {
	rapido := lista.prim
	for i := 0; i < k && rapido != nil; i++ {
		rapido = rapido.prox
	}
	if rapido == nil {
		return lista.prim.dato
	}
	lento := lista.prim
	for rapido != nil {
		rapido = rapido.prox
		lento = lento.prox
	}
	return lento.dato
}
