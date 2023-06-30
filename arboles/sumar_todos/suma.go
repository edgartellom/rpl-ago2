package suma

// Implementar una primitiva que devuelva la suma de todos los datos (números) de un árbol binario.
// Indicar y justificar el orden de la primitiva.

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) Suma() int {
	if arbol == nil {
		return 0
	}
	sumaIzq := arbol.izq.Suma()
	sumaDer := arbol.der.Suma()
	return arbol.dato + sumaIzq + sumaDer
}
