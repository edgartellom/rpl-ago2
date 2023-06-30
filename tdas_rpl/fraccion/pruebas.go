package fraccion

import (
	"fmt"
)

func EjecutarPruebas() (c, d, e Fraccion, sum_num, sum_den, prod_num, prod_den, ent int) {
	fraccion1 := CrearFraccion(10, 3)
	fmt.Printf("Fraccion 1: %d/%d\n", fraccion1.numerador, fraccion1.denominador)
	fraccion2 := CrearFraccion(5, 3)
	fmt.Printf("Fraccion 2: %d/%d\n", fraccion2.numerador, fraccion2.denominador)

	suma := fraccion1.Sumar(fraccion2)
	fmt.Printf("Suma: %d/%d\n", suma.numerador, suma.denominador)
	producto := fraccion1.Multiplicar(fraccion2)
	fmt.Printf("Suma: %d/%d\n", producto.numerador, producto.denominador)
	entero := fraccion1.ParteEntera()
	fmt.Printf("Parte Entera: %d\n", entero)

	fraccion3 := CrearFraccion(36, 48)
	fmt.Printf("Fraccion 3: %d/%d\n", fraccion3.numerador, fraccion3.denominador)
	fmt.Printf("Representacion Fraccion 3: %s\n", fraccion3.Representacion())
	fraccion4 := CrearFraccion(-16, 64)
	fmt.Printf("Fraccion 4: %d/%d\n", fraccion4.numerador, fraccion4.denominador)
	fmt.Printf("Representacion Fraccion 4: %s\n", fraccion4.Representacion())
	fraccion5 := CrearFraccion(125, 5)
	fmt.Printf("Fraccion 5: %d/%d\n", fraccion5.numerador, fraccion5.denominador)
	fmt.Printf("Representacion Fraccion 5: %s\n", fraccion5.Representacion())

	return fraccion3, fraccion4, fraccion5,
		suma.numerador, suma.denominador,
		producto.numerador, producto.denominador, entero

}
