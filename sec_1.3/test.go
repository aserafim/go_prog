package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

	fmt.Println("Entre com o seu nome: ")
	input := bufio.NewScanner(os.Stdin)
	var nome string
	for input.Scan(){
		nome = nome + input.Text()
	}
	fmt.Printf("O nome inserido foi %s", nome)
}