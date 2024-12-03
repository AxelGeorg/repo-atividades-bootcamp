package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetFileData(fileName string) string {
	defer fmt.Println("execução concluída")

	file, err := os.Open(fileName)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(data)
}

type ErrorTemValorZero struct {
	message string
}

func (e ErrorTemValorZero) Error() string {
	return "Tem valor zerado na linha"
}

type Customer struct {
	Name        string
	Id          int
	PhoneNumber int
	Adress      string
}

func verificaValorZero(custumer Customer) (Customer, error) {
	if custumer.Name == "" || custumer.Id == 0 || custumer.PhoneNumber == 0 || custumer.Adress == "" {
		return Customer{}, ErrorTemValorZero{}
	}

	return custumer, nil
}

func addCustomer(listCustomers []Customer, customer Customer) []Customer {

	err := errors.New("Error: client already exists")

	for _, cus := range listCustomers {
		if cus.Id == customer.Id {
			defer fmt.Println(err.Error())
			panic("Panic: client already exists")
		}
	}

	customerValidate, err := verificaValorZero(customer)
	if err != nil {
		defer fmt.Println("panic: Valor zerado na linha")
		panic("Panic: Valor zerado na linha")
	}

	listCustomers = append(listCustomers, customerValidate)
	return listCustomers
}

func main() {
	//Exercicio 1
	//acontece erro caso o arquivo nao exista
	txtArquivo := GetFileData("customers.txt")
	fmt.Println(txtArquivo)

	//Exercicio 2
	txtArquivo2 := GetFileData("customers.txt")
	fmt.Println(txtArquivo2)

	//
	//
	//
	//Exercicio 3
	var listCustomers []Customer

	defer fmt.Println("execução concluída")
	defer fmt.Println("End of execution")
	defer fmt.Println("Several errors were detected at runtime")

	file, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		linha := scanner.Text()
		parts := strings.Split(linha, ",")
		if len(parts) != 4 {
			fmt.Println("Linha com formato inválido:", linha)
			continue
		}

		name := strings.TrimSpace(parts[0])
		idStr := strings.TrimSpace(parts[1])
		phoneStr := strings.TrimSpace(parts[2])
		adress := strings.TrimSpace(parts[3])

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Erro ao converter id:", err)
			continue
		}

		phone, err := strconv.Atoi(phoneStr)
		if err != nil {
			fmt.Println("Erro ao converter phone:", err)
			continue
		}

		customer := Customer{
			Name:        name,
			Id:          id,
			PhoneNumber: phone,
			Adress:      adress,
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recuperado:", r)
				}
			}()

			listCustomers = addCustomer(listCustomers, customer)
		}()
	}

	fmt.Println("\nLista final de clientes:")
	for _, customer := range listCustomers {
		fmt.Printf("Nome: %s, Id: %d, Phone: R$ %d, Adress: %s\n", customer.Name, customer.Id, customer.PhoneNumber, customer.Adress)
	}
}
