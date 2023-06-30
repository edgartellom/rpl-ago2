package cambiarprioridad

type fcmpHeap[T any] func(T, T) int

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}
