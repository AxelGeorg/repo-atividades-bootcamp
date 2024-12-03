package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("dados. txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("olá, este é o meu primeiro arquivo em Go!\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	_, err = file.WriteString("Manipular arquivos é fácil com Go.\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Dados escritos no arquivo com sucesso!")

	data, err := os.ReadFile("dados. txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	fmt.Println("\nConteúdo do arquivo:")
	fmt.Println(string(data))
}
