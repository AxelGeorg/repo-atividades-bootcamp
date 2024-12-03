package main

import (
	"fmt"
	"strings"
)

const (
	opIncluir   = "1"
	opListar    = "2"
	opPesquisar = "3"
	opAlterar   = "4"
	opExluir    = "5"
	opSair      = "6"
)

type Aluno struct {
	matricula string
	nome      string
	telefone  string
	email     string
}

var listAluno []Aluno

func AdicionarAluno(aluno Aluno) {
	listAluno = append(listAluno, aluno)
}

func MostraMenu() string {
	fmt.Println("\n=== Menu de Opções ===")
	fmt.Println("1 - Incluir Aluno")
	fmt.Println("2 - Listar Todos o Alunos")
	fmt.Println("3 - Pesquisar Aluno por Matricula")
	fmt.Println("4 - Alterar Aluno")
	fmt.Println("5 - Excluir Aluno")
	fmt.Println("6 - Sair")
	fmt.Print("Escolha uma opção: ")

	var opcao string
	fmt.Scan(&opcao)
	opcao = strings.TrimSpace(opcao)

	return opcao
}

func IncluirAluno() {
	var nomeStr, telefoneStr, emailStr string

	fmt.Println("\n=== Incluir Aluno ===")

	fmt.Print("Nome: ")
	fmt.Scan(&nomeStr)

	fmt.Print("Telefone: ")
	fmt.Scan(&telefoneStr)

	fmt.Print("Email: ")
	fmt.Scan(&emailStr)

	var aluno Aluno
	aluno.nome = nomeStr
	aluno.telefone = telefoneStr
	aluno.email = emailStr

	AdicionarAluno(aluno)
}

func ListarAlunos() {
	for _, aluno := range listAluno {
		fmt.Println(aluno)
	}
}

func PesquisarAluno(matriculaStr string) {
	for _, aluno := range listAluno {
		if aluno.matricula == matriculaStr {
			fmt.Println(aluno)
			break
		}
	}

	fmt.Println("nao achou")
}

func AlterarAluno() {

}

func ExcluirAluno() {

}

func Sair() {

}

func main() {
	opcao := MostraMenu()

	for opcao != opSair {
		switch opcao {
		case opIncluir:
			IncluirAluno()
		case opListar:
			ListarAlunos()
		case opPesquisar:
			PesquisarAluno("2")
		case opAlterar:
			AlterarAluno()
		case opExluir:
			ExcluirAluno()
		case opSair:
			Sair()
		default:
			fmt.Println("\n - Essa opcao nao esta disponivel!")
		}

		opcao = MostraMenu()
	}
}
