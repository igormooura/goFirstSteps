package dataStructures

import (
	"fmt"
)

type Node struct{ 
	Data int 
	Next *Node 
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) PrintListData() {
	toPrint := l.Head
	for toPrint != nil {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
	}
	fmt.Println()
}

func (l *LinkedList) DeleteNode(value int) {
	toDelete := l.Head

	if toDelete != nil && toDelete.Data == value { // verifica se o nó a ser excluido é o primeiro (Head)
		l.Head = toDelete.Next
		l.Length--
		return
	}

	for toDelete != nil && toDelete.Next != nil {
		if toDelete.Next.Data == value {
			toDelete.Next = toDelete.Next.Next // o proximo nó vira o "depois desse", ele perde o posto de proxímo nó pro 
			l.Length--							// sucessor dele
			return
		}
		toDelete = toDelete.Next
	}
}
