package multiconj

import TDAHash "tdas/diccionario"

// Implementar el TDA MultiConjunto. Este es un Conjunto que permite más de una aparición de un elemento,
// por lo que eliminando una aparición, el elemento puede seguir perteneciendo. Dicho TDA debe tener como primitivas:

// CrearMulticonjunto[K comparable](): crea un multiconjunto.
// Guardar(elem K): guarda un elemento en el multiconjunto.
// Pertence(elem K) bool: devuelve true si el elemento aparece al menos una vez en el conjunto.
// Borrar(elem K): elimina una aparición del elemento dentro del conjunto.
// Dar la estructura del TDA y la implementación de las 4 primitivas marcadas, de forma tal que todas sean $$\mathcal{O}(1)$$.

type multiConj[K comparable] struct {
	conjunto TDAHash.Diccionario[K, TDAHash.Diccionario[K, int]]
}

func CrearMulticonjunto[K comparable]() Multiconjunto[K] {
	dicConj := TDAHash.CrearHash[K, TDAHash.Diccionario[K, int]]()
	return multiConj[K]{dicConj}
}

func (conj multiConj[K]) Guardar(elem K) {
	dicRepes := TDAHash.CrearHash[K, int]()
	if !dicRepes.Pertenece(elem) {
		dicRepes.Guardar(elem, 1)
	}
	dicRepes.Guardar(elem, dicRepes.Obtener(elem)+1)

	conj.conjunto.Guardar(elem, dicRepes)
}

func (conj multiConj[K]) Pertenece(elem K) bool {
	return conj.conjunto.Pertenece(elem)
}

func (conj multiConj[K]) Borrar(elem K) {
	if !conj.conjunto.Pertenece(elem) {
		panic("Elemento no esta en el multiconjunto")
	}
	if conj.conjunto.Obtener(elem).Obtener(elem) == 1 {
		conj.conjunto.Borrar(elem)
	} else {
		conj.conjunto.Obtener(elem).Guardar(elem, conj.conjunto.Obtener(elem).Obtener(elem)-1)
	}
}
