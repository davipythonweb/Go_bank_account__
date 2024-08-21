package contas

import (
	"Go_bank_account__/src/bank/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroConta, NumeroAgencia, Operacao int
	saldo                                float64
}

// funcao Saque que verifica se o valor eh maior
// que zero e se o valor da transaçao eh menor que o saldo.
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque de realizado com sucesso."
	} else {
		return "saldo insuficiente."
	}
}

// função para fazer deposito em conta
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.saldo
	} else {
		return "Valor do deposito invalido.", c.saldo
	}
}

// função para obter o saldo da conta
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
