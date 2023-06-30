package piramidal

import (
	"fmt"

	TDAPila "tdas/pila"
)

func EjecutarPruebas() {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 4; i >= 0; i-- {
		pila.Apilar(i)
	}
	pila2 := TDAPila.CrearPilaDinamica[int]()
	pila2.Apilar(4)
	pila2.Apilar(2)
	pila2.Apilar(3)
	fmt.Printf("La pila es piramidal: %v\n", EsPiramidal(pila))
	fmt.Printf("El tope de la pila es: %d\n", pila.VerTope())
	fmt.Printf("La pila es piramidal: %v\n", EsPiramidal(pila2))
	fmt.Printf("El tope de la pila es: %d\n", pila2.VerTope())
}
