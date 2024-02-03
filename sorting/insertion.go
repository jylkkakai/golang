package main

func insertion(arr []int) {

	temp := 0
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] > arr[i + 1] {
			temp = arr[i + 1]
			j := i
			for  {
				arr[j + 1] = arr[j]
				if j == 0 ||arr[j - 1] <= temp {
					arr[j] = temp
					break
				}
				j--
			}
		}
	}

}
