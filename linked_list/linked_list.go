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
	if list.head == nil || list.head.next == nil {
		fmt.Println("List is either empty or has only on node!")
	} else {
		current := list.head
		min := current
		for  {
			current = current.next
			if min.next == nil {
				if min.prev.data <= min.data {
					break
				}
			}
			if min.data > current.data {
				min = current
			}
			if current.next == nil {
				if min.prev != nil { 
					current = min.prev
					current.next = min.next
					if min.next != nil {
						min.next.prev = current
					}
					for current.data > min.data && current.prev != nil {
						current = current.prev
					}
					if current.prev == nil && current.data > min.data {
						min.next = list.head
						list.head = min
						min.next.prev = min
						min.prev = nil
					} else {
						min.next = current.next
						current.next = min
						min.prev = current 
						min.next.prev = min
					}
				}
				current = min.next
				min = current
				//list.printList()
			}	
		}
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
