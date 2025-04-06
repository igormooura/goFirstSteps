package pessoas

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func ShowPessoa(p *Pessoa) {
	fmt.Printf("Infos: %s , %d\n", p.Nome, p.Idade)
}
