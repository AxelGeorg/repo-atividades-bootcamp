package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	fmt.Println("\n - Lista de Produtos: ")

	for _, product := range Products {
		product.PrintProduct()
	}
}

func (p Product) PrintProduct() {
	fmt.Println("\nID: ", p.ID)
	fmt.Println("Nome: ", p.Name)
	fmt.Println("Preço: ", p.Price)
	fmt.Println("Decriçao: ", p.Description)
	fmt.Println("Categoria: ", p.Category)
}

func (p Product) GetById(id int) Product {
	for _, prod := range Products {
		if prod.ID == id {
			return prod
		}
	}

	return Product{}
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("\n - Empregado")
	fmt.Println("\n ID Empregado: ", e.ID)
	fmt.Println("Posicao: ", e.Position)
	fmt.Println("ID Pessoa: ", e.Person.ID)
	fmt.Println("Nome: ", e.Person.Name)
	fmt.Println("Data de Nascimento: ", e.Person.DateOfBirth)
}

func main() {
	//Execicio 1
	productt := Product{1, "Maça Verde", 1.8, "Maça Verde", "Fruta"}
	productt.Save()

	productt = Product{2, "Queijo", 15.5, "Queijo cheddar", "Comida"}
	productt.Save()

	productt = Product{3, "Mortadela", 18.0, "Mortadela grande", "Comida"}
	productt.Save()

	productt = Product{4, "RedBull Zero", 11.9, "RedBull Sem açucar", "Bebida"}
	productt.Save()

	productt = Product{5, "Coca Cola Zero", 13.8, "Coca Cola Sem açucar", "Bebida"}
	productt.Save()

	productt.GetAll()

	produtoBuscado := productt.GetById(3)
	produtoBuscado.PrintProduct()

	//Exercicio 2
	employee := Employee{1, "Software Developer", Person{1, "Axel", "20/05/2001"}}
	employee.PrintEmployee()
}
