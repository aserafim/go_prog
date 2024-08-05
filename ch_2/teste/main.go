package main

import (
	"fmt"
)

var test int

func main() {

	i := 0

	fmt.Printf("Variável package level: %d , variável local: %d\n", test, i)

	//Um ponteiro é o endereço da variável
	//Em outras palavras é o local aonde
	//o valor da variável é armazenado

	//Declara e define o valor de x
	x := 1

	//p recebe o endereço de memória de x
	p := &x

	//imprie o endereço de memória
	fmt.Print("\nEndereço de memória de p: ", p)

	//imprime o valor armazenado em p
	fmt.Println("\nValor armazenado em p: ", *p)

	//Altera o valor armazenado em p
	*p = 2

	//imprime novamente
	fmt.Println("Novo valor armazenado em p: ", x)

}
