package algo

import "fmt"

func merge(a1, a2 []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			res = append(res, a1[i])
			i++
		} else {
			res = append(res, a2[j])
			j++
		}
	}
	res = append(res, a1[i:]...)
	res = append(res, a2[j:]...)
	return res
}

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	mid := len(array) / 2
	left_part := MergeSort(array[:mid])
	right_part := MergeSort(array[mid:])
	return merge(left_part, right_part)
}

func runMergeSort() error {
	a1 := []int{2, 5, 8, 9, 10, 4, 3, 16, 1, 7, 8}
	fmt.Printf("Input:\n %v\n", a1)
	fmt.Printf("Output:\n %v\n", MergeSort(a1))
	return nil
}
