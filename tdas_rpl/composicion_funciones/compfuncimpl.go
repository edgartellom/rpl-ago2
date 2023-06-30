package ejercicio4

import (
	TDAPila "tdas/pila"
)

// Implementar en Go el TDA ComposiciónFunciones que emula la composición de funciones (i.e. f(g(h(x))).
//Se debe definir la estructura del TDA, y las siguientes primitivas:

// CrearComposicion() ComposicionFunciones
// AgregarFuncion(func (float64) float64)
// Aplicar(float64) float64
// Considerar que primero se irán agregando las funciones como se leen,
// pero tener en cuenta el correcto orden de aplicación.
// Por ejemplo: para emular f(g(x)), se debe hacer:

// composicion.AgregarFuncion(f)
// composicion.AgregarFuncion(g)
// composicion.Aplicar(x)
// Indicar el orden de las primitivas.

type composicionFunc struct {
	pilaFunc TDAPila.Pila[func(float64) float64]
}

func CrearComposicion() ComposicionFunciones {
	pila := TDAPila.CrearPilaDinamica[func(float64) float64]()
	return composicionFunc{pilaFunc: pila}
}

// Primitivas

func (comp composicionFunc) AgregarFuncion(funcion func(float64) float64) {
	comp.pilaFunc.Apilar(funcion)
}

func (comp composicionFunc) Aplicar(x float64) float64 {
	var backup []func(float64) float64
	for !comp.pilaFunc.EstaVacia() {
		funcion := comp.pilaFunc.Desapilar()
		x = funcion(x)
		backup = append(backup, funcion)
	}
	for i := len(backup) - 1; i >= 0; i-- {
		comp.pilaFunc.Apilar(backup[i])
	}
	return x
}
