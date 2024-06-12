package main

import "fmt"

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func buildHeap(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func pop(heap []int) ([]int,int) {
	n := len(heap)
	if n == 0 {
		return heap, -1
	}
	heap[0], heap[n-1] = heap[n-1], heap[0]
	pooped := heap[n-1]
	heap = heap[:n-1]
	heapify(heap, len(heap), 0)
	return heap, pooped

}
func main() {
	arr := []int{12, 33, 11, 5, 6, 7, 9, 20}
	fmt.Println("Original Array :", arr)
	buildHeap(arr)
	fmt.Println("Heap :", arr)
	arr, popped := pop(arr)
	if popped != -1 {
		fmt.Println("Popped element:", popped)
		fmt.Println("Max Heap after popping:", arr)
	} else {
		fmt.Println("Heap is empty.")
	}

}