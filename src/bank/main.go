package main

import (
	"Go_bank_account__/src/bank/contas"
	"fmt"
)

func main() {
	contaDaAna := contas.ContaCorrente{Titular: "Ana", Saldo: 300}
	contaDoFabio := contas.ContaCorrente{Titular: "Fabio", Saldo: 700}

	// referenciando o endereço do ponteiro com (&)
	status := contaDoFabio.Transferir(-2000, &contaDaAna)

	fmt.Println(status)
	fmt.Println(contaDaAna)
	fmt.Println(contaDoFabio)
}

// aula 10 terminada
