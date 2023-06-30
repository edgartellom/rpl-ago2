package minimo

// Implementar, por división y conquista, una función que determine el mínimo de un arreglo.
// Indicar y justificar el orden.

// BuscarMinimo devuelve el valor del minimo del arreglo, no su posicion
// Precondicion: el arreglo tiene al menos un elemento

func BuscarMinimo(arr []int) int {
	return buscarMinimo(arr, 0, len(arr)-1)
}

func buscarMinimo(arr []int, inicio, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}
	medio := (inicio + fin) / 2
	minIzq := buscarMinimo(arr, inicio, medio)
	minDer := buscarMinimo(arr, medio+1, fin)

	if minIzq < minDer {
		return minIzq
	}
	return minDer
}
