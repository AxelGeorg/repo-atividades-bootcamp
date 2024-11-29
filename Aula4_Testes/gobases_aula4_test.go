package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculaImposto(t *testing.T) {
	t.Run("Menos de 50000", func(t *testing.T) {
		salarioComImposto := calculaImposto(10000)
		expected := 0.0
		require.Equal(t, salarioComImposto, expected)
	})

	t.Run("Mais de 50000", func(t *testing.T) {
		salarioComImposto := calculaImposto(100000)
		expected := 0.17
		require.Equal(t, salarioComImposto, expected)
	})

	t.Run("Mais de 150000", func(t *testing.T) {
		salarioComImposto := calculaImposto(160000)
		expected := 0.27
		require.Equal(t, salarioComImposto, expected)
	})
}

func TestCalculaMedia(t *testing.T) {
	t.Run("Lista de notas sem nota negativa", func(t *testing.T) {
		media := calculaMedia(9, 3, 9, 3, 9)
		expected := int32(6)
		require.Equal(t, media, expected)
	})

	t.Run("Lista de notas com nota negativa", func(t *testing.T) {
		media := calculaMedia(9, 3, -9, 3, 9)
		expected := int32(6)
		require.Equal(t, media, expected)
	})
}

func TestCalculaSalario(t *testing.T) {
	t.Run("Salario categoria A", func(t *testing.T) {
		salarioCategA := calculaSalario(200, "A")
		expected := float32(10000)
		require.Equal(t, salarioCategA, expected)
	})

	t.Run("Salario categoria B", func(t *testing.T) {
		salarioCategA := calculaSalario(1200, "B")
		expected := float32(36000)
		require.Equal(t, salarioCategA, expected)
	})

	t.Run("Salario categoria C", func(t *testing.T) {
		salarioCategA := calculaSalario(1200, "C")
		expected := float32(30000)
		require.Equal(t, salarioCategA, expected)
	})

	t.Run("Salario categoria A Errado", func(t *testing.T) {
		salarioCategA := calculaSalario(5500, "A")
		expected := float32(36000)
		require.NotEqual(t, salarioCategA, expected)
	})

	t.Run("Salario categoria B Errado", func(t *testing.T) {
		salarioCategA := calculaSalario(5500, "B")
		expected := float32(10000)
		require.NotEqual(t, salarioCategA, expected)
	})

	t.Run("Salario categoria C Errado", func(t *testing.T) {
		salarioCategA := calculaSalario(5500, "C")
		expected := float32(4500)
		require.NotEqual(t, salarioCategA, expected)
	})

	t.Run("Salario categoria D, categoria errada", func(t *testing.T) {
		salarioCategA := calculaSalario(5500, "D")
		expected := float32(0)
		require.Equal(t, salarioCategA, expected)
	})
}

func TestOperation(t *testing.T) {
	t.Run("Retornou funcao minimum funcionando", func(t *testing.T) {
		minFunc, err := operation(minimum)

		require.Nil(t, err)
		require.NotNil(t, minFunc)

		min := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		expected := 1
		require.Equal(t, min, expected)
	})

	t.Run("Retornou funcao average funcionando", func(t *testing.T) {
		averageFunc, err := operation(average)

		require.Nil(t, err)
		require.NotNil(t, averageFunc)

		min := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		expected := 3
		require.Equal(t, min, expected)
	})

	t.Run("Retornou funcao maximum funcionando", func(t *testing.T) {
		maxFunc, err := operation(maximum)

		require.Nil(t, err)
		require.NotNil(t, maxFunc)

		min := maxFunc(2, 3, 3, 4, 1, 2, 17, 5)
		expected := 17
		require.Equal(t, min, expected)
	})

	t.Run("Passando operacao errada", func(t *testing.T) {
		maxFunc, err := operation("opMax")

		require.Nil(t, maxFunc)
		require.NotNil(t, err)
	})
}

func TestAnimal(t *testing.T) {
	t.Run("Teste dog retornou funcao correta", func(t *testing.T) {
		animalDog, msg := animal(dog)

		require.NotNil(t, animalDog)
		require.Equal(t, msg, "")
	})

	t.Run("Teste cat retornou funcao correta", func(t *testing.T) {
		animalCat, msg := animal(cat)

		require.NotNil(t, animalCat)
		require.Equal(t, msg, "")
	})

	t.Run("Teste hamster retornou funcao correta", func(t *testing.T) {
		animalHamster, msg := animal(hamster)

		require.NotNil(t, animalHamster)
		require.Equal(t, msg, "")
	})

	t.Run("Teste tarantula retornou funcao correta", func(t *testing.T) {
		animalTarantula, msg := animal(tarantula)

		require.NotNil(t, animalTarantula)
		require.Equal(t, msg, "")
	})

	t.Run("Teste não retorna funcao", func(t *testing.T) {
		animalDog, msg := animal("Dogg")
		require.Nil(t, animalDog)
		require.NotEqual(t, msg, "")
	})

	t.Run("Verifica soma quantidades animais", func(t *testing.T) {
		animalDog, _ := animal(dog)
		animalCat, _ := animal(cat)
		animalHamster, _ := animal(hamster)
		animalTarantula, _ := animal(tarantula)

		var amountG float64

		amountG += animalDog(10)
		amountG += animalCat(10)
		amountG += animalHamster(10)
		amountG += animalTarantula(10)

		got := amountG / 1000
		expected := float64(154)

		require.Equal(t, got, expected)
	})
}
