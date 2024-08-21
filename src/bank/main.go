package main

import (
	"Go_bank_account__/src/bank/contas"
	"fmt"
)

// funcao para pagar boleto
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

// interface para usar o metodo sacar nas duas contas
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoEliote := contas.ContaCorrente{}
	contaDaZari := contas.ContaPoupanca{}

	contaDoEliote.Depositar(5000)
	contaDaZari.Depositar(1000)

	PagarBoleto(&contaDoEliote, 1000)
	PagarBoleto(&contaDaZari, 600)

	fmt.Println(contaDoEliote.ObterSaldo())
	fmt.Println(contaDaZari.ObterSaldo())

}
