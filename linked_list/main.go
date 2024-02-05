package main

import (
	"fmt"
	"math/rand"
		)

func main () {

	fmt.Println("Hello World!")
	list := &LinkedList{}
	
	list.printList()
	list.pop()
	list.push(10)
	list.push(32)
	list.printList()
	for i := 0; i < 10; i++ {
		list.push(rand.Intn(100))
	}
	list.printList()
	fmt.Println(list)
	list.printListAll()
	list.pop()
	list.printList()
	fmt.Println(list)
	list.printListAll()
}
