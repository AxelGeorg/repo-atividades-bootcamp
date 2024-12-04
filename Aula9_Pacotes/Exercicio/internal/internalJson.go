package internal

import (
	"encoding/json"
	model "exercicio/Model"
	"io"
	"os"
)

const (
	localFileJson = "./internal/alunos.json"
)

type ServiceJson struct{}

func (s ServiceJson) LerAlunosDoArquivo() ([]model.Aluno, error) {
	var listAluno []model.Aluno

	file, err := os.Open(localFileJson)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Aluno{}, nil
		}
		return []model.Aluno{}, err
	}
	defer file.Close()

	reader := json.NewDecoder(file)

	err = reader.Decode(&listAluno)
	if err != nil {
		if err == io.EOF {
			return listAluno, nil
		}

		return nil, err
	}

	return listAluno, err
}

func (s ServiceJson) GravarAlunosEmArquivo(listAluno []model.Aluno) error {
	file, err := os.OpenFile(localFileJson, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
	}
	defer file.Close()

	writer := json.NewEncoder(file)
	return writer.Encode(listAluno)
}
