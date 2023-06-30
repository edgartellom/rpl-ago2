package ordenado

// Implementar, por división y conquista, una función que dado un arreglo y su largo, determine si el mismo se encuentra ordenado.
// Indicar y justificar el orden.

func EstaOrdenado(arr []int) bool {
	return estaOrdenado(arr, len(arr))
}

func estaOrdenado(arr []int, largo int) bool {
	if len(arr) <= 1 {
		return true
	}

	mitad := (largo) / 2

	ordenado := arr[mitad-1] <= arr[mitad]

	izq := estaOrdenado(arr[:mitad], mitad)
	der := estaOrdenado(arr[mitad:], largo-mitad)

	return izq && der && ordenado
}
