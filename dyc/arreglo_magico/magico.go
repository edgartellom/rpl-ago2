package magico

// Implementar un algoritmo en Go que reciba un arreglo de enteros de tamaño nn, ordenado ascendentemente y sin elementos repetidos,
// y determine en O(log n) si es mágico. Un arreglo es mágico si existe algún valor i tal que 0 <= i y arr[i] = i.
// Justificar el orden del algoritmo.

// Ejemplos:

// A = [ -3, 0, 1, 3, 7, 9 ] es mágico porque A[3] = 3.

// B = [ 1, 2, 4, 6, 7, 9 ] no es mágico porque B[i] != i para todo i.

func ArregloEsMagico(arr []int) bool {
	return arregloEsMagico(arr, 0, len(arr)-1)
}

func arregloEsMagico(arr []int, inicio, fin int) bool {
	if inicio > fin {
		return false
	}
	medio := (inicio + fin) / 2
	if arr[medio] == medio {
		return true
	}
	if arr[medio] < medio {
		return arregloEsMagico(arr, medio+1, fin)
	} else {
		return arregloEsMagico(arr, inicio, medio-1)
	}
}
