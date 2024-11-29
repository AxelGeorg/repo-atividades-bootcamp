package main

import "fmt"

func main() {
	// Exercício 1
	var nome, endereco string

	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)

	fmt.Print("Digite seu endereço: ") // Corrigido aqui
	fmt.Scan(&endereco)

	fmt.Printf("Olá, %s! Seu endereço é %s.\n", nome, endereco) // Corrigido para refletir o endereço

	// Exercício 2
	var temperatura float64 = 22.5
	var umidade int = 60
	var pressao float64 = 1013.25

	fmt.Printf("Temperatura: %.2f°C\n", temperatura)
	fmt.Printf("Umidade: %d%%\n", umidade)
	fmt.Printf("Pressão: %.2f hPa\n", pressao)

	// Exercício 3
	var firstName string
	var lastName string
	var age int
	lastName = "Georg" // Corrigido para um sobrenome mais lógico
	var driver_license bool = true
	var person int
	childsNumber := 2

	// Exibindo as variáveis
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Age:", age)
	fmt.Println("Driver License:", driver_license)
	fmt.Println("Person:", person)
	fmt.Println("Childs Number:", childsNumber)

	// Exercício 4
	var lastName2 string = "Smith"
	var age2 int = 35
	boolean := false
	var salary float32 = 45857.90
	var firstName3 string = "Mary"

	// Exibindo as variáveis
	fmt.Println("Last Name 2:", lastName2)
	fmt.Println("Age 2:", age2)
	fmt.Println("Boolean:", boolean)
	fmt.Println("Salary:", salary)
	fmt.Println("First Name 3:", firstName3)
}

/*
O que é um pacote em Golang e qual a sua importância em um programa?
R: Um pacote é uma forma de agrupamento que fica dentro do modulo, cada pacote pode ter um conjunto de arquivos com extesão .go

Como declarar e utilizar o pacote principal (main) em uma aplicação Go?
R:Declarar: Package main.. para utilizar basta criar uma funcao main que faça algo e rodar um go run e o nome do arquivo com a func main.

Qual é a diferença entre pacotes internos (standard library) e pacotes personalizados em Go?
R:

Como importar e utilizar um pacote externo no Go? (Dê um exemplo.)
R:

Qual é o comando utilizado no Go para inicializar um novo módulo e gerenciar dependências?
R:go mod init modulo

O que acontece se você importar um pacote no Go e não utilizá-lo no código?
R: Some da linha de imports ou da erro

Como declarar uma variável global em um pacote e torná-la acessível a outros pacotes?
R: basta criar a variavel fora de uma func e o nome dela precisa iniciar com um caractere maiusculo

Qual a diferença entre uma variável declarada com var e uma variável inicializada com := em Go?
R: na utiliza de var a tipagem fica declarada de forma explicita e com := é a inferencia, assim o tipo mais adequado é atribuido de forma automatica

Como você exporta uma variável ou função de um pacote em Golang, e qual a convenção de nomenclatura para isso?
R: Utilizar a primeira letra maiuscula sendo na variavel ou na funcao

Explique o conceito de shadowing (sombras) em relação às variáveis no Go e como isso pode afetar o escopo.
R:

*/
