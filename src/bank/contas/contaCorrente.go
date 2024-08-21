package contas

import (
	"Go_bank_account__/src/bank/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64 //letra-minuscula para deixar o saldo visivel somente aqui
}

// funcao Saque que verifica se o valor eh maior
// que zero e se o valor da transaçao eh menor que o saldo.
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "saldo insuficiente."
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

// função para obter o saldo da conta
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
