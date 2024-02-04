package main

func mergeCombine(larr, rarr []int) []int {

	l := 0
	r := 0 
	m := 0
	merged_arr := make([]int, len(larr) + len(rarr))
	for m < len(merged_arr) {
		if l > len(larr) - 1 {
			merged_arr[m] = rarr[r]
			m++
			r++
		} else if r > len(rarr) - 1 { 
			merged_arr[m] = larr[l]
			m++
			l++
		} else if larr[l] <= rarr[r] {
			merged_arr[m] = larr[l]
			m++
			l++
		} else {
			merged_arr[m] = rarr[r]
			m++
			r++
		}

	}
	return merged_arr
}

func merge(arr []int) []int {
	
	if len(arr) == 1 {
		return arr
	}
	a := merge(arr[0:len(arr)/2])
	b := merge(arr[len(arr)/2:])


	return mergeCombine(a, b)
}

