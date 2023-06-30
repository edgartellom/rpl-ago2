package magico

import (
	"fmt"
)

func EjecutarPruebas() {
	arr1 := []int{-3, 0, 1, 3, 7, 9}
	arr2 := []int{1, 2, 4, 6, 7, 9}
	fmt.Println(ArregloEsMagico(arr1))
	fmt.Println(ArregloEsMagico(arr2))
}
