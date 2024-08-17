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

	// outra forma de criar struct
	// o (*) eh para usar um ponteiro para fazer referencia a algo na memoria
	var contaDaZari *ContaCorrente
	contaDaZari = new(ContaCorrente)
	contaDaZari.titular = "Zari"
	contaDaZari.saldo = 350
	fmt.Println(*contaDaZari)
}
