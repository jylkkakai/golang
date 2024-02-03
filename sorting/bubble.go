package main


func bubble (arr []int) {

	left := 0
	right := 0
	temp := 0
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < (len(arr) - 1 - i); j++ {
		
			left = arr[j]
			right = arr[j + 1]

			if left > right {
				temp = left 
				arr[j] = right
				arr[j + 1] = temp
			}
		}
	}

}
