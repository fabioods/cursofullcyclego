package main

import (
	"cursofullcyclego/error"
	"cursofullcyclego/function"
	"cursofullcyclego/math"
	"cursofullcyclego/ponteiro"
	"cursofullcyclego/structs"
	"encoding/json"
	"fmt"
)

type Carro struct {
	Nome string
}

func (c Carro) Andar() {
	fmt.Printf("\n\nO carro %v andou", c.Nome)
}

func modificaValor(valor *int) int {
	*valor += 1
	return *valor
}

func main() {
	fmt.Print("hello world \n")
	a := 10
	b := 3.14
	c := "hello"
	d := `
		falaaaa
		devvvv
	`
	e := true
	fmt.Printf("%T - %v \n", a, a)
	fmt.Printf("%T - %v \n", b, b)
	fmt.Printf("%T - %v \n", c, c)
	fmt.Printf("%T - %v \n", d, d)
	fmt.Printf("%T - %v \n", e, e)

	fmt.Println(math.SomaX(1))
	// fmt.Println(math.somaX(1))
	fmt.Println(math.A)
	error.TestError()
	fmt.Println(error.SomaError(1, 11))
	somaFunctionResult := function.SomaFunc(1, 11)
	fmt.Println(somaFunctionResult)

	somaVaridic := function.SomaVariadic(1, 2, 4, 5, 6, 7, 9, 10)
	fmt.Printf("soma variadica %v", somaVaridic)

	resultFuntion := func(a int, b int) func() int {
		return func() int {
			return a + b
		}
	}
	fmt.Printf("\nfunction result %v", resultFuntion(1, 2)())

	carro := Carro{
		Nome: "Gol",
	}
	carro.Andar()

	ponteiro.Ponteiro()

	value := 10
	fmt.Printf("\n value before %v", value) //10
	valorAlterado := modificaValor(&value)
	fmt.Printf("\n value after %v", value) // 11
	fmt.Printf("\n valorAlterado %v", valorAlterado)

	fmt.Printf("\n\n cliente %v", structs.DeclareCliente().Nome)
	fmt.Printf("\n\n other cliente %v", structs.OtherDeclareCliente().Nome)
	fmt.Printf("\n\n cliente internacional %v", structs.DeclareClienteInternacional())

	clienteJson, err := json.Marshal(structs.DeclareClienteInternacional())
	clienteJsonString := string(clienteJson)
	fmt.Printf("\n\n cliente json %v", clienteJsonString)
	fmt.Printf("\n\n cliente json err %v", err)

	clienteHidratado := structs.ClienteInternacional{}
	fmt.Printf("\n\n cliente hidratado before %v", clienteHidratado)
	json.Unmarshal([]byte(clienteJsonString), &clienteHidratado)
	fmt.Printf("\n\n cliente hidratado after %v", clienteHidratado)

}
