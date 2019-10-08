package algo

import "fmt"

func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	pivot, i := array[0], 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[i] > pivot {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			i++
			head++
		}
	}
	//array[0], array[i] = array[i], array[0]
	QuickSort(array[:head])
	QuickSort(array[head+1:])
}

func runQuickSort() error {
	a1 := []int{2, 5, 8, 9, 10, 4, 3, 16, 1, 7, 8}
	//a1 := []int{3, 4, 1, 2}
	fmt.Printf("Input:\n %v\n", a1)
	QuickSort(a1)
	fmt.Printf("Output:\n %v\n", a1)
	return nil
}
