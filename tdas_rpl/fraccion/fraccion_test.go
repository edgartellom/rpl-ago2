package fraccion_test

import (
	"tdas_rpl/fraccion"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFraccion(t *testing.T) {
	t.Log("Pruebas de TDA Fraccion")
	fraccion3, fraccion4, fraccion5, suma_num, suma_den, product_num, product_den, entero := fraccion.EjecutarPruebas()

	require.EqualValues(t, 45, suma_num)
	require.EqualValues(t, 9, suma_den)
	require.EqualValues(t, 50, product_num)
	require.EqualValues(t, 9, product_den)
	require.EqualValues(t, 3, entero)

	require.EqualValues(t, "3/4", fraccion3.Representacion())
	require.EqualValues(t, "-1/4", fraccion4.Representacion())
	require.EqualValues(t, "25", fraccion5.Representacion())

}
