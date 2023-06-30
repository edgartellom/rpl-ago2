package pilalargo

// Implementar una función recursiva que reciba una pila y devuelva la cantidad de elementos de la misma.
//  Al terminar la ejecución de la función la pila debe quedar en el mismo estado al original.

func Largo[T any](pila Pila[T]) int {
	if pila.EstaVacia() {
		return 0
	}

	elemento := pila.Desapilar()
	largo := Largo(pila) + 1
	pila.Apilar(elemento)
	return largo
}
