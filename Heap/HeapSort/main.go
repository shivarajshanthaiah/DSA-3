package main

import "fmt"

// This us max heap and will sort array in
// ascending order and min heap will sort
// in descending order...
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

func heapSort(arr []int) {
	// building heap
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}

	//sorting
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr[:i], 0)
	}
}

func main() {
	arr := []int{12, 43, 35, 22, 66, 23, 31, 7, 5, 24, 53, 6}
	fmt.Println(arr)

	heapSort(arr)
	fmt.Println(arr)

}
