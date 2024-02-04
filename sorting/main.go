package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRndArr(arr []int) {

	for i := 0; i < len(arr); i++ {
//		fmt.Println(rand.Intn(100))
		arr[i] = rand.Intn(len(arr))
	}
	
}

func timeSort(size int, iter int) {

	rnd_arr := make([]int, size)
	//start := 0
	fmt.Printf("Running sort algorithms for %d times...\n", iter)
	fmt.Printf("Array size: %d\n", size)
	fmt.Println("Bubble")
	elapsed_time := 0.0
	for i := 0; i < iter; i++ {
		createRndArr(rnd_arr)
		start := time.Now()
		bubble(rnd_arr)
		elapsed_time += float64(time.Since(start))
	}
	fmt.Printf("Search time: %f ms\n", elapsed_time/1000000/float64(iter))
	
	fmt.Println("Selection")
	elapsed_time = 0
	for i := 0; i < iter; i++ {
		createRndArr(rnd_arr)
		start := time.Now()
		selection(rnd_arr)
		elapsed_time += float64(time.Since(start))
	}
	fmt.Printf("Search time: %f ms\n", elapsed_time/1000000/float64(iter))

	fmt.Println("Insertion")
	elapsed_time = 0
	for i := 0; i < iter; i++ {
		createRndArr(rnd_arr)
		start := time.Now()
		insertion(rnd_arr)
		elapsed_time += float64(time.Since(start))
	}
	fmt.Printf("Search time: %f ms\n", elapsed_time/1000000/float64(iter))

	fmt.Println("Merge")
	elapsed_time = 0
	for i := 0; i < iter; i++ {
		createRndArr(rnd_arr)
		start := time.Now()
		merge(rnd_arr)
		elapsed_time += float64(time.Since(start))
	}
	fmt.Printf("Search time: %f ms\n", elapsed_time/1000000/float64(iter))

	fmt.Println("Quick")
	elapsed_time = 0
	for i := 0; i < iter; i++ {
		createRndArr(rnd_arr)
		start := time.Now()
		quick(rnd_arr, 0, size - 1)
		elapsed_time += float64(time.Since(start))
	}
	fmt.Printf("Search time: %f ms\n", elapsed_time/1000000/float64(iter))

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

	fmt.Println("Quick")
	createRndArr(rnd_arr)
	fmt.Println(rnd_arr)
	quick(rnd_arr, 0, size - 1)
	fmt.Println(rnd_arr)

	timeSort(10, 1000)
	timeSort(100, 1000)
	timeSort(1000, 1000)
	timeSort(10000, 100)
	timeSort(100000, 100)




}


