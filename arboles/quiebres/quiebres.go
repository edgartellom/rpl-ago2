package quiebres

// Definimos como quiebre en un árbol binario cuando ocurre que:

// un hijo derecho tiene un solo hijo, y es el izquierdo
// un hijo izquierdo tiene un solo hijo, y es el derecho
// Implementar una primitiva para el árbol binario size_t ab_quiebres(const ab_t*) que, dado un árbol binario,
// nos devuelva la cantidad de quiebres que tiene. La primitiva no debe modificar el árbol.

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) Quiebres() int {
	if arbol == nil {
		return 0
	}
	izq := arbol.izq.Quiebres()
	der := arbol.der.Quiebres()

	suma := izq + der
	if arbol.der != nil && arbol.der.tieneUnSoloHijo() && arbol.der.izq != nil {
		suma += 1
	}
	if arbol.izq != nil && arbol.izq.tieneUnSoloHijo() && arbol.izq.der != nil {
		suma += 1
	}
	return suma
}

func (arbol *ab) tieneUnSoloHijo() bool {
	if arbol == nil {
		return false
	}
	if arbol.der != nil && arbol.izq != nil {
		return false
	}
	return arbol.der != nil || arbol.izq != nil
}
