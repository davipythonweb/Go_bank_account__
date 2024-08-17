package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoEliote := ContaCorrente{titular: "Eliote", numeroAgencia: 589,
		numeroConta: 11525, saldo: 125.5}

	contaDaMariah := ContaCorrente{"Mariah", 222, 24321, 200}

	fmt.Println(contaDoEliote)
	fmt.Println(contaDaMariah)
}
