package main

import (
	"Go_bank_account__/src/bank/clientes"
	"Go_bank_account__/src/bank/contas"
	"fmt"
)

func main() {

	// Fazendo de outra forma
	clienteEliote := clientes.Titular{"Eliote", "123.123.123.12", "Desenvolvedor-Go"}
	contaDoEliote := contas.ContaCorrente{clienteEliote, 0034, 11525, 300}

	// contaDoEliote := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "Eliote",
	// 	Cpf:       "123.123.123.12",
	// 	Profissao: "Desenvolvedor-Full-Stack"},

	// 	NumeroAgencia: 0034, NumeroConta: 111525, Saldo: 300}

	fmt.Println(contaDoEliote)

}

// aula 12 finalizada
