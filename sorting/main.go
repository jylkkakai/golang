package main

import (
	"fmt"
	"math/rand"
)

func createRndArr(arr []int) {

	for i := 0; i < len(arr); i++ {
//		fmt.Println(rand.Intn(100))
		arr[i] = rand.Intn(len(arr))
	}
	
}

func main() {
	fmt.Println("Hello World")
	size := 20 
	rnd_arr := make([]int, size) 

	fmt.Println("Bubble")
	createRndArr(rnd_arr)
	fmt.Println(rnd_arr)
	bubble(rnd_arr)
	fmt.Println(rnd_arr)

	fmt.Println("Selection")
	createRndArr(rnd_arr)
	fmt.Println(rnd_arr)
	selection(rnd_arr)
	fmt.Println(rnd_arr)

	fmt.Println("Insertion")
	createRndArr(rnd_arr)
	fmt.Println(rnd_arr)
	insertion(rnd_arr)
	fmt.Println(rnd_arr)

	fmt.Println("Merge")
	createRndArr(rnd_arr)
	fmt.Println(rnd_arr)
	marr := merge(rnd_arr)
	fmt.Println(marr)

}


