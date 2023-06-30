package multiprimeros

// Implementar la primitiva func (cola *colaEnlazada[T]) Multiprimeros(k int) []T que dada una cola y un número k,
// devuelva los primeros k elementos de la cola, en el mismo orden en el que habrían salido de la cola.
// En caso que la cola tenga menos de k elementos.
// Si hay menos elementos que k en la cola, devolver un slice del tamaño de la cola.
// Indicar y justificar el orden de ejecución del algoritmo.

func (cola *colaEnlazada[T]) Multiprimeros(k int) []T {
	var resul []T
	cont := 0
	for actual := cola.primero; actual != nil && cont < k; actual = actual.siguiente {
		resul = append(resul, actual.dato)
		cont++
	}
	return resul
}
