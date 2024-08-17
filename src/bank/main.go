package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// funcao Saque que verifica se o valor eh maior
// que zero e se o valor da transaçao eh menor que o saldo.
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque de realizado com sucesso.", c.saldo
	} else {
		return "Saldo insuficiente.", c.saldo
	}
}

// função para fazer deposito em conta
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.saldo
	} else {
		return "Valor do deposito invalido.", c.saldo
	}
}

// funçao para fazer a transferencia
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaAna := ContaCorrente{titular: "Ana", saldo: 300}
	contaDoFabio := ContaCorrente{titular: "Fabio", saldo: 700}

	// referenciando o endereço do ponteiro com (&)
	status := contaDoFabio.Transferir(-2000, &contaDaAna)

	fmt.Println(status)
	fmt.Println(contaDaAna)
	fmt.Println(contaDoFabio)
}

// aula 9 terminada
