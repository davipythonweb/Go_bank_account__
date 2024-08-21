package main

import (
	"Go_bank_account__/src/bank/contas"
	"fmt"
)

func main() {

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())

}
