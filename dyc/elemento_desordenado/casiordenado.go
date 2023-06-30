package casiordenado

// Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos
// y casi ordenado (todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar.
// Indicar y justificar el orden.

func ElementoDesordenado(arr []int) int {
	return elementoDesordenado(arr, 0, len(arr)-1)
}

func elementoDesordenado(arr []int, inicio, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}

	medio := (inicio + fin) / 2

	if arr[medio] > arr[medio+1] {
		return elementoDesordenado(arr, inicio, medio)
	}
	return elementoDesordenado(arr, medio+1, fin)
}
