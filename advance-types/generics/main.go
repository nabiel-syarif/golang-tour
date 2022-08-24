package main

import "fmt"

func Print[T any](data T) {
	fmt.Println(data)
}

type LinkedList[T any] struct {
	Val  T
	Next *LinkedList[T]
}

func (list *LinkedList[T]) Print() {
	curr := list
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func main() {
	Print("Halo")

	list := LinkedList[int]{
		Val: 1,
		Next: &LinkedList[int]{
			Val: 2,
		},
	}
	list.Next.Next = &LinkedList[int]{
		Val: 3,
	}
	list.Print()
}
