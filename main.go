package main

import (
	"fmt"

	// TDAS
	// FiltrarCola "tdas_rpl/cola_filtrar"
	// TDACompos "tdas_rpl/composicion_funciones"
	// Balanceado "tdas_rpl/expresion_balanceada"
	// TDAFraccion "tdas_rpl/fraccion"
	// MergePilas "tdas_rpl/merge_pilas"
	// OrdenaPila "tdas_rpl/ordenar_pila"
	// PilaLargo "tdas_rpl/pila_largo"
	// Piramidal "tdas_rpl/pila_piramidal"
	// SumaPares "tdas_rpl/visitar_lista_pares"

	// DYC
	EsMagico "dyc/arreglo_magico"
	Ordenado "dyc/arreglo_ordenado"
	Minimo "dyc/buscar_minimo"
	Desordenado "dyc/elemento_desordenado"
)

func main() {
	// TDAFraccion.EjecutarPruebas()
	// fmt.Println("*************************")
	// TDACompos.EjecutarPruebas()
	// fmt.Println("*************************")
	// Piramidal.EjecutarPruebas()
	// fmt.Println("*************************")
	// OrdenaPila.EjecutarPruebas()
	// fmt.Println("*************************")
	// FiltrarCola.EjecutarPruebas()
	// fmt.Println("\n*************************")
	// SumaPares.EjecutarPruebas()
	// fmt.Println("*************************")
	// MergePilas.EjecutarPruebas()
	// fmt.Println("*************************")
	// PilaLargo.EjecutarPruebas()
	// fmt.Println("*************************")
	// Balanceado.EjecutarPruebas()
	EsMagico.EjecutarPruebas()
	fmt.Println("*************************")
	Minimo.EjecutarPruebas()
	fmt.Println("*************************")
	Ordenado.EjecutarPruebas()
	fmt.Println("*************************")
	Desordenado.EjecutarPruebas()
}
