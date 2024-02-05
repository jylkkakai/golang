package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) push(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		newNode.prev = current
		current.next = newNode
		// fmt.Println(current.next)
	}
}

func (list *LinkedList) pop () {
	if list.head == nil {
		fmt.Println("Empty list!")
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current = current.prev
		current.next = nil
	}
}

func (list *LinkedList) sort() {
	if list.head == nil {
		fmt.Println("Empty list!")
	}

}

func (list *LinkedList) printList() {
	if list.head == nil {
		fmt.Println("Empty list.")
	} else {
		current := list.head
		for current.next != nil {
			fmt.Print(current.data, " ")
			current = current.next
		}
		fmt.Print(current.data, "\n")
	}
}


func (list *LinkedList) printListAll() {
	if list.head == nil {
		fmt.Println("Empty list.")
	} else {
		current := list.head
		for current.next != nil {
			fmt.Println(current)
			current = current.next
		}
		fmt.Println(current)
	}
}
