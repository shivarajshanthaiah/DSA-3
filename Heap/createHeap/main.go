package main

import "fmt"

// this is max heap & heap down
// for min heap change condition to '<'
func heapify(arr []int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < len(arr) && arr[left] > arr[largest] {
		largest = left
	}
	if right < len(arr) && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest)
	}
}

// in max and min heap largest & smallest values
// will be their root value respectively
func peek(arr []int) int {
	return arr[0]
}

func buildHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}
}

func main() {
	arr := []int{12, 33, 11, 5, 6, 7, 9, 20}
	fmt.Println("origina arry :", arr)
	buildHeap(arr)
	fmt.Println("Heapified array :", arr)

	fmt.Println(peek(arr))
}
