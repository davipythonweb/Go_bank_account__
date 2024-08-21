package main

import (
	"Go_bank_account__/src/bank/contas"
	"fmt"
)

func main() {

	contaDoEliote := contas.ContaCorrente{}
	contaDaZari := contas.ContaPoupanca{}

	contaDaZari.Depositar(1000)
	contaDaZari.Sacar(500)

	fmt.Println(contaDoEliote)
	fmt.Println(contaDaZari.ObterSaldo())

}
