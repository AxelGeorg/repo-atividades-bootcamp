package main

import "fmt"

type Student struct {
	id      int
	name    string
	surname string
	date    string
}

var Students []Student

func (s Student) Save() {
	Students = append(Students, s)
}

func (s Student) Print() {
	fmt.Println("\nID: ", s.id)
	fmt.Println("Nome: ", s.name)
	fmt.Println("Sobrenome: ", s.surname)
	fmt.Println("Data: ", s.date)
}

func (s Student) PrintAll() {
	fmt.Println("\n - Lista de Estudantes: ")

	for _, student := range Students {
		student.Print()
	}
}

type Product interface {
	Price() float64
}

type SmallProtuct struct {
	price float64
}

func (s SmallProtuct) Price() float64 {
	return s.price
}

type AvarageProtuct struct {
	price float64
}

func (a AvarageProtuct) Price() float64 {
	return a.price + (a.price * 0.06)
}

type BigProtuct struct {
	price float64
}

func (b BigProtuct) Price() float64 {
	return b.price + (b.price * 0.06) + 2500
}

const (
	small   = 1
	average = 2
	big     = 3
)

func Factory(typeProd int, pricee float64) Product {
	switch typeProd {
	case small:
		return SmallProtuct{pricee}
	case average:
		return AvarageProtuct{pricee}
	case big:
		return BigProtuct{pricee}
	default:
		return nil
	}
}

func main() {
	// Exercicio 1
	student1 := Student{1, "Axel", "Georg", "20/05/2001"}
	student1.Save()
	student1.Print()

	student2 := Student{2, "Felipe", "Matheus", "27/12/2001"}
	student2.Save()
	student2.Print()

	student3 := Student{3, "Thiago", "Coppi", "23/09/1995"}
	student3.Save()
	student3.Print()

	student3.PrintAll()

	//Exercicio 2
	smallProduct := Factory(small, 1000)
	avarageProduct := Factory(average, 1500)
	bigProduct := Factory(big, 3000)

	fmt.Println("\n\n Produtos: ")
	fmt.Println("Preco do produto pequenno é", smallProduct.Price())
	fmt.Println("Preco do produto medio é", avarageProduct.Price())
	fmt.Println("Preco do produto grande é", bigProduct.Price())
}
