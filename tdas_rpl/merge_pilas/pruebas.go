package mergepilas

import (
	"fmt"

	TDAPila "tdas/pila"
)

func EjecutarPruebas() {
	pila1 := TDAPila.CrearPilaDinamica[int]()
	pila2 := TDAPila.CrearPilaDinamica[int]()

	pila1.Apilar(1)
	pila1.Apilar(3)
	pila1.Apilar(3)
	pila1.Apilar(5)

	pila2.Apilar(2)
	pila2.Apilar(4)
	pila2.Apilar(5)
	pila2.Apilar(6)
	pila2.Apilar(6)

	fmt.Println(MergePilas(pila1, pila2))
}
