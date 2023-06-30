package sumapares

import (
	"fmt"

	TDALista "tdas/lista"
)

func EjecutarPruebas() {
	lista := TDALista.CrearListaEnlazada[*int]()
	arr := []*int{new(int), new(int), new(int), new(int), new(int)}
	for i := 0; i < 5; i++ {
		*arr[i] = i
		lista.InsertarUltimo(arr[i])
	}
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		fmt.Println(*iter.VerActual())
	}
	fmt.Println("Suma de Pares:", SumaPares(lista))
}
