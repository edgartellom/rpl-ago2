package cambiarprioridad

// Implementar en Go una primitiva para el heap (siendo este un max-heap) que reciba un heap
// y una función de comparación y lo reordene de manera tal que se se comporte como max-heap
// para la nueva función de comparación (se cambia la función de prioridad).
// El orden de dicha primitiva debe ser O(n).

func (heap *heap[T]) CambiarPrioridad(nuevaPrioridad func(a, b T) int) {

}
