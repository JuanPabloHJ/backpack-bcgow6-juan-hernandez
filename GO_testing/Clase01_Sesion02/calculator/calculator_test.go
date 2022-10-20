package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestarTesting(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 12
	num2 := 4
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion Restar() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}
}

func TestRestarTestify(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 12
	num2 := 4
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestSort(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	input := []int{9, 8, 7, 6, 6, 5, 4, 3, 2, 1}
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9}

	// Se ejecuta el test
	resultado := SortAscending(input)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividirError(t *testing.T) {
	num1 := 5
	num2 := 0

	_, resultadoEsperado := Dividir(num1, num2)

	assert.NotNil(t, resultadoEsperado, resultadoEsperado.Error())
}

func TestDividir(t *testing.T) {
	num1 := 6
	num2 := 2

	resultadoEsperado := 3
	resultado, _ := Dividir(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, fmt.Sprintf("los resultados no coinciden"))
}
