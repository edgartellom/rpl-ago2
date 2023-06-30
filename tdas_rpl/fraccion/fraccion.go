package fraccion

import "fmt"

// Implementar el TDA Fracción.
// Dicho TDA debe tener las siguientes primitivas, cuya documentación puede encontrarse en fraccion.go:

// CrearFraccion(numerador, denominador int) Fraccion
// Sumar(otra Fraccion) Fraccion
// Multiplicar(otra Fraccion) Fraccion
// ParteEntera() int
// Representacion() string
// Nota: considerar que se puede utilizar la función del módulo fmt Sprintf para generar la
// representación de la fracción.

type Fraccion struct {
	numerador   int
	denominador int
}

// CrearFraccion crea una fraccion con el numerador y denominador indicados.
// Si el denominador es 0, entra en panico.
func CrearFraccion(numerador, denominador int) Fraccion {
	fraccion := new(Fraccion)
	fraccion.numerador = numerador
	if denominador == 0 {
		panic("Denominador no puede ser cero")
	}
	fraccion.denominador = denominador
	return *fraccion
}

// Sumar crea una nueva fraccion, con el resultante de hacer la suma de las fracciones originales
func (f Fraccion) Sumar(otra Fraccion) Fraccion {
	suma := new(Fraccion)
	suma.denominador = f.denominador * otra.denominador
	suma.numerador = otra.denominador*f.numerador + f.denominador*otra.numerador
	return *suma
}

// Multiplicar crea una nueva fraccion con el resultante de multiplicar ambas fracciones originales
func (f Fraccion) Multiplicar(otra Fraccion) Fraccion {
	producto := new(Fraccion)
	producto.numerador = f.numerador * otra.numerador
	producto.denominador = f.denominador * otra.denominador
	return *producto
}

// ParteEntera devuelve la parte entera del numero representado por la fracción.
// Por ejemplo, para "7/2" = 3.5 debe devolver 3.
func (f Fraccion) ParteEntera() int {
	return f.numerador / f.denominador
}

// Representacion devuelve una representación en cadena de la fraccion simplificada (por ejemplo, no puede devolverse
// "10/8" sino que debe ser "5/4"). Considerar que si se trata de un número entero, debe mostrarse como tal.
// Considerar tambien el caso que se trate de un número negativo.
func (f Fraccion) Representacion() string {

	d := mcd(f.numerador, f.denominador)
	numerador := f.numerador / d
	denominador := f.denominador / d

	if denominador == 1 {
		return fmt.Sprintf("%d", numerador)
	}
	if denominador < 0 {
		numerador = -numerador
		denominador = -denominador
	}

	return fmt.Sprintf("%d/%d", numerador, denominador)
}

func mcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
