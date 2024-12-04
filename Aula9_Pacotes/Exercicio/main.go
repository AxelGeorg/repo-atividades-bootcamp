package main

import (
	"errors"
	model "exercicio/Model"
	"exercicio/internal"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

const (
	opIncluir   = "1"
	opListar    = "2"
	opPesquisar = "3"
	opAlterar   = "4"
	opExluir    = "5"
	opSair      = "6"
)

var listAluno []model.Aluno
var serviceJson internal.ServiceJson
var servicetxt internal.ServiceTxt

func GravarInformacoes() error {
	if err := serviceJson.GravarAlunosEmArquivo(listAluno); err != nil {
		return err
	}

	if err := servicetxt.GravarAlunosEmArquivo(listAluno); err != nil {
		return err
	}

	return nil
}

func LerInformacoes() ([]model.Aluno, error) {

	var err error
	var listaAlunosJson []model.Aluno
	var listaAlunosTxt []model.Aluno

	if listaAlunosJson, err = serviceJson.LerAlunosDoArquivo(); err != nil {
		return []model.Aluno{}, err
	}

	if listaAlunosTxt, err = servicetxt.LerAlunosDoArquivo(); err != nil {
		return []model.Aluno{}, err
	}

	if len(listaAlunosJson) != len(listaAlunosTxt) {
		return []model.Aluno{}, errors.New("As lista no json e no txt estao com quantidades diferentes!")
	}

	return listaAlunosJson, nil
}

func AdicionarAluno(aluno model.Aluno) error {
	listAluno = append(listAluno, aluno)

	fmt.Println("Aluno incluido com sucesso!!")
	return GravarInformacoes()
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

func IncluirAluno() error {
	var nomeStr, telefoneStr, emailStr string

	fmt.Println("\n=== Incluir Aluno ===")

	fmt.Print("Nome: ")
	fmt.Scan(&nomeStr)

	fmt.Print("Telefone: ")
	fmt.Scan(&telefoneStr)

	fmt.Print("Email: ")
	fmt.Scan(&emailStr)

	aluno := model.Aluno{
		Matricula: uuid.New().String(),
		Nome:      nomeStr,
		Telefone:  telefoneStr,
		Email:     emailStr,
	}

	return AdicionarAluno(aluno)
}

func MostraAluno(aluno model.Aluno) {
	fmt.Println("Matricula:", aluno.Matricula, "- Nome:", aluno.Nome, "- Email:", aluno.Email, "- Telefone:", aluno.Telefone)
}

func ListarAlunos() {
	fmt.Println("\n -- Lista de Alunos --")

	for _, aluno := range listAluno {
		MostraAluno(aluno)
	}
}

func MenuBuscarAlunoPorMatricula() (string, error) {
	var matriculaStr string
	fmt.Print("Digite a Matricula:")
	fmt.Scan(&matriculaStr)

	if strings.TrimSpace(matriculaStr) == "" {
		return "", errors.New("A Matricula nao foi digitada!")
	}

	return matriculaStr, nil
}

func BuscarAlunoPorMatricula(matriculaStr string) (model.Aluno, error) {
	for _, aluno := range listAluno {
		if aluno.Matricula == matriculaStr {
			return aluno, nil
		}
	}

	return model.Aluno{}, errors.New("Nao existe aluno com esta Matricula!!")
}

func AlterarInfosAluno(aluno *model.Aluno) bool {
	var nomeStr, telefoneStr, emailStr string
	var alunoAlterado bool

	fmt.Printf("Novo Nome (Antigo: %s): ", aluno.Nome)
	fmt.Scanln(&nomeStr)
	if nomeStr == "" {
		nomeStr = aluno.Nome
	} else {
		alunoAlterado = true
	}

	fmt.Printf("Novo Telefone (Antigo: %s): ", aluno.Telefone)
	fmt.Scanln(&telefoneStr)
	if telefoneStr == "" {
		telefoneStr = aluno.Telefone
	} else {
		alunoAlterado = true
	}

	fmt.Printf("Novo Email (Antigo: %s): ", aluno.Email)
	fmt.Scanln(&emailStr)
	if emailStr == "" {
		emailStr = aluno.Email
	} else {
		alunoAlterado = true
	}

	aluno.Nome = nomeStr
	aluno.Telefone = telefoneStr
	aluno.Email = emailStr

	return alunoAlterado
}

func PesquisarAluno() {
	fmt.Println("\n=== Pesquisar Aluno ===")

	Matricula, err := MenuBuscarAlunoPorMatricula()
	if err != nil {
		fmt.Println(err)
	}

	aluno, err := BuscarAlunoPorMatricula(Matricula)
	if err != nil {
		fmt.Println(err)
	}

	MostraAluno(aluno)
}

func AlterarAluno() {
	fmt.Println("\n=== Alterar Aluno ===")

	matricula, err := MenuBuscarAlunoPorMatricula()
	if err != nil {
		fmt.Println(err)
	}

	aluno, err := BuscarAlunoPorMatricula(matricula)
	if err != nil {
		fmt.Println(err)
	}

	alterou := AlterarInfosAluno(&aluno)

	if alterou {
		for i, a := range listAluno {
			if a.Matricula == matricula {
				listAluno[i] = aluno
				break
			}
		}
		if err := GravarInformacoes(); err != nil {
			fmt.Println("Erro ao gravar no arquivo:", err)
		}
		fmt.Println("Aluno alterado com sucesso!!")
	} else {
		fmt.Println("Nenhuma alteração foi feita.")
	}

	MostraAluno(aluno)
}

func ExcluirAluno() {
	fmt.Println("\n=== Excluir Aluno ===")

	Matricula, err := MenuBuscarAlunoPorMatricula()
	if err != nil {
		fmt.Println(err)
		return
	}

	index := -1
	for i, aluno := range listAluno {
		if aluno.Matricula == Matricula {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Nao existe aluno com esta Matricula!")
		return
	}

	listAluno = append(listAluno[:index], listAluno[index+1:]...)
	if err := GravarInformacoes(); err != nil {
		fmt.Println("Erro ao gravar no arquivo:", err)
	}
	fmt.Println("Aluno excluido com sucesso!")
}

func Sair() {
	fmt.Println("Saindo do sistema...")
	os.Exit(0)
}

func main() {
	var err error
	if listAluno, err = serviceJson.LerAlunosDoArquivo(); err != nil {
		fmt.Println("Erro ao ler alunos do arquivo:", err)
		return
	}

	opcao := MostraMenu()

	for opcao != opSair {
		switch opcao {
		case opIncluir:
			IncluirAluno()
		case opListar:
			ListarAlunos()
		case opPesquisar:
			PesquisarAluno()
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
