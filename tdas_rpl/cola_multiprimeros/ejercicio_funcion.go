package multiprimeros

// Implementar la función func Multiprimeros[T any](cola Cola[T], k int) []T con el mismo comportamiento de la primitiva.

func Multiprimeros[T any](cola Cola[T], k int) []T {
	var resul []T
	for !cola.EstaVacia() {
		resul = append(resul, cola.Desencolar())
	}
	for _, elem := range resul {
		cola.Encolar(elem)
	}
	if k > len(resul) {
		return resul
	}
	return resul[:k]
}
