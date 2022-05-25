package ellipsis

import "fmt"

func One(valores ...int64) any {
	for _, valor := range valores {
		fmt.Println(valor)
	}
	return nil
}

func GetAlunos(alunos ...string) any {
	for _, aluno := range alunos {
		fmt.Println("Olá Aluno(a)", aluno)
	}
	return nil
}
