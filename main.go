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

}
