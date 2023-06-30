package pilalargo

import (
	"fmt"

	TDAPila "tdas/pila"
)

func EjecutarPruebas() {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 100; i++ {
		pila.Apilar(i)
	}
	fmt.Println(Largo[int](pila))
}
