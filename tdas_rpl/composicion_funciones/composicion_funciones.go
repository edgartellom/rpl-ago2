package ejercicio4

type ComposicionFunciones interface {
	AgregarFuncion(func(float64) float64)
	Aplicar(float64) float64
}
