package main

import (
	"awesomeProject/math"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// := -> declaração e atribuição com inferência de tipo
	a := "cleiton"
	b := true
	c := 32
	d := 1.1
	e := ` lalal
	teste `

	fmt.Println("aqui o nome:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	result := math.Soma(1, 1)
	fmt.Println("soma: ", result)

	result2 := math.SquareX(2000000000)
	fmt.Println(result2)

	// array com make

	var tamanho int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scanln(&tamanho)

	arr := make([]int, tamanho) // make já pré determina

	fmt.Println("Digite os números do array:")
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&arr[i])
	}

	fmt.Println("\nArray digitado:", arr)

	// array com append

	var tamanho2 int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scanln(&tamanho2)

	var arr2 []int // slice vazio

	fmt.Println("Digite os números do array:")
	for i := 0; i < tamanho2; i++ {
		var valor int
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&valor)
		arr2 = append(arr2, valor)
	}

	fmt.Println("\nArray digitado:", arr2)

}
