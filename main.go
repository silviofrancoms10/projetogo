package main

import "fmt"

func main() {

	//Ponteiros

	x := 10
	y := &x

	fmt.Println(x, *y) // quando coloco asterisco na frente do Y estou me referindo ao VALOR dela
	fmt.Println(&x, y) // quando coloco o & na frente do X estou me referindo ao endereço dela, neste caso o Y já é o endereço de X

	testandoPonteiros(&x, y)
	fmt.Println(x, *y)
	fmt.Println(&x, y)

	// Exercicio 1
	/*arr := [2]int{6, 3}

	total := arr[0] + arr[1]

	fmt.Println(total)

	arr2 := [3]int{3, 8, 9}

	var total2 int

	for i := 0; i < len(arr2); i++ {
		total2 += arr2[i]
	}

	fmt.Println(total2)*/

	// Exercicio 2
	/*lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	var sublistaMenor int
	var sublistaMaior int

	for i := 0; i < len(lista); i++ {
		if lista[i] < 6 {
			sublistaMenor += lista[i]
		} else {
			sublistaMaior += lista[i]

		}

	}
	fmt.Println(sublistaMenor, sublistaMaior)*/

}

func testandoPonteiros(x *int, y *int) {
	*x = 20
}
