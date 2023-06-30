package ejercicio4

import (
	"fmt"
)

func EjecutarPruebas() {
	func1 := func(x float64) float64 { return x * x }
	func2 := func(x float64) float64 { return x + 5 }
	func3 := func(x float64) float64 { return x * 3 }

	comp := CrearComposicion()
	comp.AgregarFuncion(func3)
	comp.AgregarFuncion(func2)
	comp.AgregarFuncion(func1)
	res := comp.Aplicar(3)
	fmt.Printf("Resultado: %v\n", res)
	res2 := comp.Aplicar(res)
	fmt.Printf("Resultado: %v\n", res2)
}
