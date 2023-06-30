package sumapares

import (
	TDALista "tdas/lista"
)

// Sabiendo que la firma del iterador interno de la lista enlazada es:

//     Iterar(visitar func(K) bool)
// Se tiene una lista en donde todos los elementos son punteros a números enteros.
// Implementar una función SumaPares que reciba una lista y, utilizando el iterador interno (no el externo),
// calcule la suma de todos los números pares.

func SumaPares(lista TDALista.Lista[*int]) int {
	suma := 0
	sumaPtr := &suma
	lista.Iterar(func(i *int) bool {
		if (*i)%2 == 0 {
			*sumaPtr += *i
		}
		return true
	})

	return *sumaPtr
}
