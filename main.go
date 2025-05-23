package main

import (
	"awesomeProject/dataStructures"
	"awesomeProject/math"
	"awesomeProject/pessoas"
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

	fmt.Println("================================")
	// array com append

	var tamanho3 int

	fmt.Printf("Digite aqui o tamanho do array3: ")
	fmt.Scanln(&tamanho3)

	var arr3 []int

	fmt.Printf("Digite aqui os elementos do array3: ")

	for i := 0; i < tamanho3; i++ {
		var x int
		fmt.Scanln(&x)
		arr3 = append(arr3, x)
	}

	fmt.Println(arr3)

	fmt.Println("================================")

	cleiton := pessoas.Pessoa{
		Nome:  "Cleiton Silva",
		Idade: 33,
	}

	pessoas.ShowPessoa(&cleiton)

	cleiton.Nome = "Cleitinho mudou de nome"

	pessoas.ShowPessoa(&cleiton)

	fmt.Println("================================")

	//  Linked List

	ll := &dataStructures.LinkedList{}

	ll.Prepend(&dataStructures.Node{Data: 10})
	ll.Prepend(&dataStructures.Node{Data: 20})
	ll.Prepend(&dataStructures.Node{Data: 30})

	fmt.Println("Lista após adicionar elementos:")
	ll.PrintListData()

	ll.DeleteNode(20)

	fmt.Println("Lista após deletar o nó com valor 20:")
	ll.PrintListData()

	ll.DeleteNode(30)
	ll.DeleteNode(10)

	fmt.Println("Lista após deletar o nó com valor 30 e 10:")
	ll.PrintListData()

	// stack

	fmt.Println("Stack: ")

	stack := &dataStructures.Stack{}

	stack.Push(10)
	stack.Push(42)
	stack.Print()

	stack.Pop()
	stack.Print()

	fmt.Println("================================")
	fmt.Println("Binary Search Tree Operations:")

	tree := dataStructures.NewTree(10)
	tree = dataStructures.Insert(tree, 5)
	tree = dataStructures.Insert(tree, 15)
	tree = dataStructures.Insert(tree, 2)
	tree = dataStructures.Insert(tree, 7)
	tree = dataStructures.Insert(tree, 12)
	tree = dataStructures.Insert(tree, 18)

	fmt.Println("Initial tree:")
	dataStructures.PrintTree(tree, 0)

	var values []int
	dataStructures.TraverseInOrder(tree, &values)
	fmt.Println("\nValues in order:", values)

	fmt.Println("\nDeleting value 5:")
	tree = dataStructures.Delete(tree, 5)
	dataStructures.PrintTree(tree, 0)

	fmt.Println("\nInserting value 20:")
	tree = dataStructures.Insert(tree, 20)
	dataStructures.PrintTree(tree, 0)

	values = []int{}
	dataStructures.TraverseInOrder(tree, &values)
	fmt.Println("\nFinal values in order:", values)

}
