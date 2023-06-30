package balanceada

import (
	"fmt"
)

func EjecutarPruebas() {
	fmt.Println(Balanceado("[{([])}]"))
	fmt.Println(Balanceado("[{}"))
	fmt.Println(Balanceado("[(])"))
	fmt.Println(Balanceado("()[{}]"))
	fmt.Println(Balanceado("()()(())"))
}
