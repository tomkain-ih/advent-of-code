package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StackInstruction struct {
	Quantity    int
	From        int
	Destination int
}

func (s *StackInstruction) String() string {
	return fmt.Sprint(s.Quantity, s.From, s.Destination)
}

func (s *StackInstruction) NewStackInstruction(raw string) {
	strs := strings.Split(raw, " ")
	s.Quantity, _ = strconv.Atoi(strs[1])
	s.From, _ = strconv.Atoi(strs[3])
	s.Destination, _ = strconv.Atoi(strs[5])
}

func (s *StackInstruction) Execute(stacks []*Stack) {
	for i := 0; i < s.Quantity; i++ {
		v := stacks[s.From-1].Pop()
		stacks[s.Destination-1].Push(v)
	}
}

func (s *StackInstruction) ExecuteMulti(stacks []*Stack) {
	queue := NewStack()
	for i := 0; i < s.Quantity; i++ {
		v := stacks[s.From-1].Pop()
		queue.Push(v)
	}
	for range queue.nodes {
		pop := queue.Pop()
		stacks[s.Destination-1].Push(pop)
	}
	// ^ this worked
	/*for i := 0; i < queue.count; i++ {
		pop := queue.Pop()
		fmt.Println(charValue(pop.Value))
		stacks[s.Destination-1].Push(pop)
	}*/
	// ^ this didn't
}
