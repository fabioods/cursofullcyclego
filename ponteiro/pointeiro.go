package ponteiro

import "fmt"

func Ponteiro() {
	a := 10
	var ponteiro *int = &a
	b := &a
	fmt.Printf("\n Ponteiro de a %v", &a)
	fmt.Printf("\n Valor de a %v", a)
	fmt.Printf("\n Ponteiro de b %v", b)
	fmt.Printf("\n Valor de b %v", *b)
	fmt.Printf("\n Ponteiro de ponteiro %v",ponteiro)
	fmt.Printf("\n Valor de ponteiro %v",*ponteiro)
	*ponteiro = 50
	fmt.Printf("\n Valor da variável ponteiro foi alterado para 50")
	fmt.Printf("\n Valor de a atualizado %v", a)
	fmt.Printf("\n Valor de b atualizado %v", *b)
}

