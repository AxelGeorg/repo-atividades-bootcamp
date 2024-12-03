package main

import (
	"errors"
	"fmt"
)

var SalaryError = errors.New("Error: the salary entered does not reach the taxable minimum")

type SalaryErrorS struct {
	message string
}

func (e SalaryErrorS) Error() string {
	return e.message
}

func SalarioMensal(horasTrab float32, valorHora float32) (float32, error) {

	salario := horasTrab * valorHora
	if salario >= 150000 {
		salario -= salario * 0.1
	}

	if horasTrab < 80.0 {
		return 0.0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	return salario, nil
}

func main() {
	//Exercicio 1
	salary1 := 140000
	salary2 := 167000

	if salary1 < 150000 {
		fmt.Println(SalaryError.Error())
	} else {
		fmt.Println("Must pay tax")
	}

	if salary2 < 150000 {
		fmt.Println(SalaryError.Error())
	} else {
		fmt.Println("Must pay tax")
	}

	//Exercicio 2
	salary3 := 99000

	if salary3 <= 100000 {
		err1 := SalaryErrorS{message: "Error: salary is less than 10000"}

		if errors.Is(err1, SalaryErrorS{message: "Error: salary is less than 10000"}) {
			fmt.Println(err1.message)
		}
	} else {
		fmt.Println("Salary is sufficient")
	}

	//Exercicio 3
	salary4 := 50000

	if salary4 <= 100000 {

		var err2 = errors.New("Error: the salary entered does not reach the taxable minimum")
		erroSalaryy := err2
		if errors.Is(err2, erroSalaryy) {
			fmt.Println(err2.Error())
		}
	} else {
		fmt.Println("Salary is sufficient")
	}

	//Exercicio 4
	salary5 := 58000

	if salary5 <= 100000 {
		err3 := fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary5)
		fmt.Println(err3.Error())
	}

	//Exercicio 5
	salario6, err6 := SalarioMensal(79, 100.0)
	if err6 != nil {
		fmt.Println(err6.Error())
	} else {
		fmt.Println("Salario ok", salario6)
	}
}
