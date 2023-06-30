package altura

// Dado un árbol binario, escribir una primitiva recursiva que determine la altura del mismo.
// Indicar y justificar el orden de la primitiva.

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) Altura() int {
	if arbol == nil {
		return 0
	}
	alturaIzq := arbol.izq.Altura()
	alturaDer := arbol.der.Altura()
	if alturaIzq > alturaDer {
		return alturaIzq + 1
	}
	return alturaDer + 1
}
