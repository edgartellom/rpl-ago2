package balanceada

import TDAPila "tdas/pila"

// Implementar una función func Balanceado(texto string) bool, que retorne si texto esta balanceado o no.
// texto sólo puede contener los siguientes caracteres: [,],{,}(,).
// Indicar y justificar la complejidad de la función implementada.
// Un texto esta balanceado si cada agrupador abre y cierra en un orden correcto. Por ejemplo:

// balanceado("[{([])}]") => true
// balanceado("[{}") => false
// balanceado("[(])") => false
// balanceado("()[{}]") => true
// balanceado("()()(())") => true

func Balanceado(texto string) bool {
	pila := TDAPila.CrearPilaDinamica[rune]()
	for _, char := range texto {
		switch char {
		case '(', '[', '{':
			pila.Apilar(char)
		case ')', ']', '}':
			if pila.EstaVacia() {
				return false
			}
			elem := pila.Desapilar()
			if !esParAgrupador(elem, char) {
				return false
			}
		}
	}
	return pila.EstaVacia()
}

func esParAgrupador(apertura, cierre rune) bool {
	switch apertura {
	case '{':
		return cierre == '}'
	case '[':
		return cierre == ']'
	case '(':
		return cierre == ')'
	default:
		return false
	}
}
