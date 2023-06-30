package reconstruccion

import "strings"

// Determinar cómo es el Árbol cuyo pre order es EURMAONDVSZT, e in order es MRAUOZSVDNET.

type ab struct {
	izq   *ab
	der   *ab
	clave string
}

func ReconstruirParticular() *ab {
	preorder := "EURMAONDVSZT"
	inorder := "MRAUOZSVDNET"
	preorderArr := strings.Split(preorder, "")
	inorderArr := strings.Split(inorder, "")
	return Reconstruir(preorderArr, inorderArr)
}

func Reconstruir(preorder, inorder []string) *ab {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	raiz := preorder[0]
	ab := &ab{clave: raiz}
	posRaiz := buscarPosRaiz(raiz, inorder)

	inorderIzq := inorder[:posRaiz]
	inorderDer := inorder[posRaiz+1:]
	preorderIzq := preorder[1 : len(inorderIzq)+1]
	preorderDer := preorder[len(inorderIzq)+1:]

	ab.izq = Reconstruir(preorderIzq, inorderIzq)
	ab.der = Reconstruir(preorderDer, inorderDer)
	return ab
}

func buscarPosRaiz(raiz string, inorder []string) int {
	pos := 0
	for pos := 0; pos < len(inorder); pos++ {
		if inorder[pos] == raiz {
			return pos
		}
	}
	return pos
}
