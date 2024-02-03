package main
func selection(arr []int) {

	min_index := 0
	temp := 0
	for i := 0; i < len(arr); i++ {
		min_index = i
		for j := 0 + 1 + i; j < len(arr); j++ {
			if arr[min_index] > arr[j] {
				min_index = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min_index]
		arr[min_index] = temp
	}
}
