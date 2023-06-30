package doshijos

// Dado un árbol binario, escriba una primitiva recursiva que cuente la cantidad de nodos que tienen exactamente dos hijos directos.
// ¿Qué orden de complejidad tiene la función implementada?

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) DosHijos() int {
	if arbol == nil {
		return 0
	}
	cont := 0
	if arbol.izq != nil && arbol.der != nil {
		cont = 1
	}
	cont += arbol.izq.DosHijos()
	cont += arbol.der.DosHijos()
	return cont
}
