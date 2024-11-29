package main

import "fmt"

func calculaImposto(salario float64) float64 {

	var porcentagemImposto float64
	switch {
	case salario > 150000:
		porcentagemImposto += 0.10
		fallthrough
	case salario > 50000:
		porcentagemImposto += 0.17
	}

	return porcentagemImposto
}

func calculaMedia(notas ...int32) int32 {
	var somaNotas, qtdNotasValidas int32
	for _, nota := range notas {
		if nota < 0 {
			fmt.Println("Uma nota negativa foi encontrada e ignorada no calculo!")
			continue
		}

		qtdNotasValidas++
		somaNotas += nota
	}

	return somaNotas / qtdNotasValidas
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(op string) (func(...int) int, error) {
	switch op {
	case minimum:
		return func(numbers ...int) int {
			min := numbers[0]
			for _, number := range numbers {
				if number < min {
					min = number
				}
			}

			return min
		}, nil
	case average:
		return func(numbers ...int) int {
			var soma int
			for _, number := range numbers {
				soma += number
			}

			return soma / len(numbers)
		}, nil
	case maximum:
		return func(numbers ...int) int {
			max := numbers[0]
			for _, number := range numbers {
				if number > max {
					max = number
				}
			}

			return max
		}, nil
	default:
		return nil, fmt.Errorf("invalid operation: %s", op)
	}
}

func calculaSalario(tempoMin uint, categ string) float32 {
	tempoHoras := float32(tempoMin) / 60

	switch categ {
	case "A":
		return tempoHoras * 3000
	case "B":
		salario := tempoHoras * 1500
		return salario + (salario * 0.2)
	case "C":
		salario := tempoHoras * 1000
		return salario + (salario * 0.5)
	}

	return 0
}

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animal(animal string) (func(quant int) float64, string) {
	switch animal {
	case cat:
		return func(quant int) float64 {
			return float64(quant) * (10 * 1000)
		}, ""
	case dog:
		return func(quant int) float64 {
			return float64(quant) * (5 * 1000)
		}, ""
	case hamster:
		return func(quant int) float64 {
			return float64(quant) * 250
		}, ""
	case tarantula:
		return func(quant int) float64 {
			return float64(quant) * 150
		}, ""
	default:
		return nil, "Animal invalido!"
	}
}

func main() {
	//Exercicio 1
	var salario float64

	fmt.Print("Digite o seu salario: ")
	fmt.Scan(&salario)

	fmt.Println("O Imposto é", calculaImposto(salario)*salario)

	//Exercicio 2
	mediaNotas := calculaMedia(9, 3, -9, 3, 9)
	fmt.Println("A media das notas é", mediaNotas)

	//Exercicio 3
	salarioCategA, salarioCategB, salarioCategC := calculaSalario(9600, "A"), calculaSalario(9600, "B"), calculaSalario(9600, "C")
	fmt.Println("Salario Funcionario Categoria A é ", salarioCategA)
	fmt.Println("Salario Funcionario Categoria B é ", salarioCategB)
	fmt.Println("Salario Funcionario Categoria C é ", salarioCategC)

	//Exercicio 4
	minFunc, err := operation(minimum)

	if err != nil {
		fmt.Println(err)
		return
	}

	averageFunc, err := operation(average)

	if err != nil {
		fmt.Println(err)
		return
	}

	maxFunc, err := operation(maximum)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Minimo:", minFunc(2, 3, 3, 4, 10, 2, 4, 5))
	fmt.Println("Média:", averageFunc(2, 3, 3, 4, 1, 2, 4, 5))
	fmt.Println("Máximo:", maxFunc(2, 3, 3, 4, 1, 2, 4, 5))

	//Exercicio 5
	animalDog, msg := animal(dog)

	if msg != "" {
		fmt.Println(msg)
		return
	}

	animalCat, msg := animal(cat)

	if msg != "" {
		fmt.Println(msg)
		return
	}

	animalHamster, msg := animal(hamster)

	if msg != "" {
		fmt.Println(msg)
		return
	}

	animalTarantula, msg := animal(tarantula)

	if msg != "" {
		fmt.Println(msg)
		return
	}

	var amountG float64

	amountG += animalDog(10)
	amountG += animalCat(10)
	amountG += animalHamster(10)
	amountG += animalTarantula(10)

	fmt.Println("Total animais:", amountG/1000)

}
