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

func main() {
	contaDaAna := ContaCorrente{}
	contaDaAna.titular = "Ana"
	contaDaAna.saldo = 500.00

	fmt.Println(contaDaAna.saldo)

	fmt.Println(contaDaAna.Sacar(100))

	fmt.Println(contaDaAna.saldo)

	fmt.Println(contaDaAna.Depositar(-2000))
}

// func main() {
// 	contaDoEliote := ContaCorrente{titular: "Eliote", numeroAgencia: 589,
// 		numeroConta: 11525, saldo: 125.5}

// 	contaDaMariah := ContaCorrente{"Mariah", 222, 24321, 200}

// 	fmt.Println(contaDoEliote)
// 	fmt.Println(contaDaMariah)

// 	// outra forma de criar struct
// 	// o (*) eh para usar um ponteiro para fazer referencia a algo na memoria
// 	var contaDaZari *ContaCorrente
// 	contaDaZari = new(ContaCorrente)
// 	contaDaZari.titular = "Zari"
// 	contaDaZari.saldo = 350
// 	fmt.Println(*contaDaZari)

// 	// uso do (&) para saber o endereço de fato na memoria
// 	fmt.Println(&contaDaZari)

// }
