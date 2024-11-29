package main

import "fmt"

func main() {
	//Execicio 1
	/*
		var palavra string

		fmt.Print("Digite sua palavra: ")
		fmt.Scan(&palavra)

		for _, char := range palavra {
			fmt.Println(string(char))
		}

		//Execicio 2
		var idade, tempoEmpregadoAtual int
		var salario float32

		fmt.Print("Digite sua idade: ")
		fmt.Scan(&idade)

		fmt.Print("Digite a quantidade de anos que esta na empresa atual: ")
		fmt.Scan(&tempoEmpregadoAtual)

		fmt.Print("Digite seu salario: ")
		fmt.Scan(&salario)

		if idade <= 22 {
			fmt.Print("O banco só concede empréstimos a clientes com mais de 22 anos de idade")
		} else if tempoEmpregadoAtual < 1 {
			fmt.Print("O banco só concede empréstimos a clientes que estejam empregados")
		} else if tempoEmpregadoAtual <= 1 {
			fmt.Print("O banco só concede empréstimos a clientes que estejam em seu emprego há mais de um ano")
		}

		if salario <= 100000 {
			fmt.Print("Voce pode receber o emprestivo, mas dentro dos empréstimos concedidos, o banco não cobrará juros daqueles que tiverem um salário superior a US$ 100.000")
		}
	*/
	//Execicio 3
	var numMes int
	var mes string

	fmt.Print("Digite o numero do mes: ")
	fmt.Scan(&numMes)

	switch numMes {
	case 1:
		mes = "Janeiro"
	case 2:
		mes = "Fevereiro"
	case 3:
		mes = "Marco"
	case 4:
		mes = "April"
	case 5:
		mes = "Maio"
	case 6:
		mes = "Junho"
	case 7:
		mes = "Julho"
	case 8:
		mes = "Agosto"
	case 9:
		mes = "Setembro"
	case 10:
		mes = "Outubro"
	case 11:
		mes = "Novembro"
	case 12:
		mes = "Dezembro"
	default:
		mes = "Indefinido"
	}

	fmt.Println("O mes é", mes)

	var meses = map[int]string{
		1: "Janeiro", 2: "Fevereiro", 3: "Março", 4: "Abril", 5: "Maio", 6: "Junho",
		7: "Julho", 8: "Agosto", 9: "Setembro", 10: "Outubro", 11: "Novembro", 12: "Dezembro",
	}

	if mesNome, mesValido := meses[numMes]; mesValido {
		fmt.Println("O mês é", mesNome)
	} else {
		fmt.Println("O mês é Indefinido")
	}

	// Exercício 4
	var employees = map[string]int{
		"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30,
	}

	fmt.Printf("Idade do Benjamin é %d\n", employees["Benjamin"])

	var qtdFunc int

	for _, idadee := range employees {
		if idadee > 21 {
			qtdFunc++
		}
	}

	fmt.Printf("Quantidade de funcionários com mais de 21 anos é %d\n", qtdFunc)

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)
}
