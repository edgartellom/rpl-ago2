package ordenarpila

import (
	"fmt"
	TDAPila "tdas/pila"
)

func EjecutarPruebas() {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(4)
	pila.Apilar(1)
	pila.Apilar(5)
	pila.Apilar(2)
	pila.Apilar(3)
	Ordenar(pila)
	fmt.Println(pila)
}
