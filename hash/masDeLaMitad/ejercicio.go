package masmitad

import TDAHash "tdas/diccionario"

// Implementar una función de orden O(n) que dado un arreglo de n números enteros
//  devuelva true o false según si existe algún elemento que aparezca más de la mitad de las veces.
// Justificar el orden de la solución. Ejemplos:

// [1, 2, 1, 2, 3] -> false
// [1, 1, 2, 3] -> false
// [1, 2, 3, 1, 1, 1] -> true
// [1] -> true

func MasDeLaMitad(arr []int) bool {
	mitad := len(arr) / 2
	dic := TDAHash.CrearHash[int, int]()
	for _, numero := range arr {
		if !dic.Pertenece(numero) {
			dic.Guardar(numero, 0)
		}
		dic.Guardar(numero, dic.Obtener(numero)+1)
		if dic.Obtener(numero) > mitad {
			return true
		}
	}
	return false
}
