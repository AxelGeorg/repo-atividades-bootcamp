package internal

import (
	"bufio"
	model "exercicio/Model"
	"fmt"
	"os"
	"strings"
)

const (
	localFiletxt = "./internal/alunos.txt"
)

type ServiceTxt struct{}

func (s ServiceTxt) LerAlunosDoArquivo() ([]model.Aluno, error) {
	file, err := os.Open(localFiletxt)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Aluno{}, nil
		}
		return []model.Aluno{}, err
	}
	defer file.Close()

	var alunos []model.Aluno
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) == 4 {
			alunos = append(alunos, model.Aluno{
				Matricula: parts[0],
				Nome:      parts[1],
				Telefone:  parts[2],
				Email:     parts[3],
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return []model.Aluno{}, err
	}

	return alunos, nil
}

func (s ServiceTxt) GravarAlunosEmArquivo(listAluno []model.Aluno) error {
	file, err := os.OpenFile(localFiletxt, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, aluno := range listAluno {
		line := fmt.Sprintf("%s;%s;%s;%s\n", aluno.Matricula, aluno.Nome, aluno.Telefone, aluno.Email)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
