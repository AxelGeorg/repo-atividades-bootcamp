package model

type Aluno struct {
	Matricula string `json:"matricula"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
}
