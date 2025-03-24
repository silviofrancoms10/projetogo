package main

import "fmt"

func main() {
	//aqui classe é chamada de struct

	e := endereco{
		rua:    "Rua XV",
		numero: 973,
		cidade: "Campo Grande",
	}

	fmt.Println(e)

}

type endereco struct {
	rua    string
	numero int
	cidade string
}
