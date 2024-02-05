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
	list.sort()
	list.push(10)
	list.sort()
	list.push(32)
	list.printList()
	for i := 0; i < 10; i++ {
		list.push(rand.Intn(100))
	}
	list.printList()
	//list.printListAll()
	list.pop()
	list.printList()
	list.sort()
	list.printList()
}
