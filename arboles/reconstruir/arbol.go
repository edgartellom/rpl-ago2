package reconstruir

// Implementar una primitiva para el AB que reciba dos arreglos (o listas) de cadenas.
// El primer arreglo corresponde al preorder de un árbol binario.
// El segundo al inorder del mismo árbol (ambos arreglos tienen los mismos elementos, sin repetidos).
// La función debe devolver un árbol binario que tenga dicho preorder e inorder.
// Indicar y justificar el orden de la primitiva (tener cuidado con este punto).
// Considerar que la estructura del árbol binario es:

type Arbol struct {
	izq   *Arbol
	der   *Arbol
	clave int
}

func Reconstruir(preorder, inorder []int) *Arbol {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	raiz := preorder[0]
	ab := &Arbol{clave: raiz}
	posRaiz := buscarPosRaiz(raiz, inorder)

	inorderIzq := inorder[:posRaiz]
	inorderDer := inorder[posRaiz+1:]
	preorderIzq := preorder[1 : len(inorderIzq)+1]
	preorderDer := preorder[len(inorderIzq)+1:]

	ab.izq = Reconstruir(preorderIzq, inorderIzq)
	ab.der = Reconstruir(preorderDer, inorderDer)
	return ab
}

func buscarPosRaiz(raiz int, inorder []int) int {
	pos := 0
	for pos := 0; pos < len(inorder); pos++ {
		if inorder[pos] == raiz {
			return pos
		}
	}
	return pos
}
