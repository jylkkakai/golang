package main

func quick(arr []int, lo int, hi int) {

	if lo > hi {
		return
	}
	pivot := arr[hi]
	temp := 0
	j := 0
	for i := range arr {

		if arr[i] < pivot {
			temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			j++
		}
		if i == hi {
			arr[i] = arr[j]
			arr[j] = pivot
		}
	}
	//quick(arr[:j], lo, j - 1)
	//quick(arr[j + 1:], j + 1, hi)
	quick(arr[:j], 0, j - 1)
	quick(arr[j + 1:], 0, hi - j - 1)

}
