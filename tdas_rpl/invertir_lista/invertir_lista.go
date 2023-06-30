package invertir

// Implementar en Go una primitiva func (lista *listaEnlazada[T]) Invertir() que invierta la lista,
// sin utilizar estructuras auxiliares. Indicar y justificar el orden de la primitiva.

func (lista *listaEnlazada[T]) Invertir() {
	var arr []T
	for actual := lista.primero; actual != nil; actual = actual.siguiente {
		arr = append(arr, actual.dato)
	}
	i := lista.largo - 1
	for actual := lista.primero; actual != nil; actual = actual.siguiente {
		actual.dato = arr[i]
		i--
	}
}
