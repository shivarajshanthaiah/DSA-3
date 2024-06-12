package main

import "fmt"

// func heapifyUp(arr []int, i int) {
// 	parent := (i - 1) / 2

// 	for i > 0 && arr[parent] < arr[i] {
// 		arr[parent], arr[i] = arr[i], arr[parent]
// 		i = parent
// 		parent = (i - 1) / 2
// 	}
// }

func heapifyDown(arr []int, i int) {
	parent := (i - 1) / 2

	for i > 0 && arr[parent] > arr[i] {
		arr[parent], arr[i] = arr[i], arr[parent]
		i = parent
		parent = (i - 1) / 2
	}
}

func insertIntoHeap(arr []int, val int) []int {
	arr = append(arr, val)
	heapifyDown(arr, len(arr)-1)
	return arr
}

func main() {

	arr := []int{33, 20, 11, 5, 6, 7, 9, 12}
	fmt.Println("Original", arr)

	newVal := 25
	res := insertIntoHeap(arr, newVal)
	fmt.Println("After insertion", res)
}
