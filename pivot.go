package main

import "fmt"

func pivot() {
	arr := []int{2, 1, 4, 3, 6}
	fmt.Println("soucre arr :", arr)
	h := len(arr) - 1
	a, p := getPivot(arr, 0, h)
	fmt.Println("arr :", a, ",pivot :", p)
}

func getPivot(arr []int, l, h int) ([]int, int) {

	p := arr[h]
	i := l

	for j := l; j < h; j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[h] = arr[h], arr[i]
	return arr, p
}
