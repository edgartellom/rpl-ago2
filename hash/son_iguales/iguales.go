package iguales

import TDAHash "tdas/diccionario"

// Asumiendo que se tiene disponible una implementación completa del TDA Hash,
// se desea implementar una función que decida si dos Hash dados representan o no el mismo Diccionario.
// Considere para la solución que es de interés la mejor eficiencia temporal posible.
// Indique, para su solución, eficiencia en tiempo y espacio.
// Nota: Dos tablas de hash representan el mismo diccionario si tienen la misma cantidad de elementos;
// todas las claves del primero están en el segundo; todas las del segundo, en el primero;
// y los datos asociados a cada una de esas claves son iguales (se pueden comparar los valores con “==”).

func SonIguales[K comparable, V comparable](d1, d2 TDAHash.Diccionario[K, V]) bool {
	if d1.Cantidad() != d2.Cantidad() {
		return false
	}
	for iter := d1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !d2.Pertenece(clave) || d2.Obtener(clave) != valor {
			return false
		}
	}
	return true
}
