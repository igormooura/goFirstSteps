package dataStructures

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
